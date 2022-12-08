package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Scans the directory for files with the specified extension and uploads them to S3 bucket.
// Also takes a timeout value to terminate the update if it doesn't complete within that time.
//
// The AWS credentials and Region should be provided as environmental variables, see README for the details.
//
// Usage:
//   # Scan the directory for '.md' files and upload them to S3. Must complete within 10 minutes or will fail
//   go run main.go -t 10m -d docs -e .md

//   # Scan the current directory for '.tar.gz' files and upload them to S3. Must complete within 3 minutes or will fail
//   go run main.go

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil { return e }
		fileName := d.Name()
		if len(fileName) > len(ext) {
			lastChars := fileName[len(fileName)-len(ext):]
			if lastChars == ext {
				a = append(a, s)
			}
		}
		return nil
	})
	return a
}

func main() {
	var directory, extension string
	var timeout time.Duration

	bucket := os.Getenv("AWS_BUCKET")
	fmt.Println("Bucket is", bucket)

	flag.StringVar(&directory, "d", ".", "Directory name.")
	flag.StringVar(&extension, "e", ".tar.gz", "Extension.")
	flag.DurationVar(&timeout, "t", 60*3*time.Second, "Upload timeout.")
	flag.Parse()

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Find all files and upload them to S3
	for _, artifact := range find(directory, extension) {
		fmt.Printf("found a file to upload: %s\n", artifact)
		f, err  := os.Open(artifact)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file %q, %v", artifact, err)
		}

		// Upload the file to S3.
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(artifact),
			Body:   f,
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to upload file, %v", err)
		} else {
			fmt.Printf("file uploaded to %s\n", aws.StringValue(&result.Location))
			os.Setenv("RESULTS_PATH", artifact)
		}
	}
}

