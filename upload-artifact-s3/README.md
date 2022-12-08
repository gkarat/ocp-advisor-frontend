upload-artifact-s3
------------------

A simple tool to scan the directory recursively and upload files of the specified extension to AWS S3 bucket.


Configuration
-------------

The following environmental variables should be present:
```shell
$ export AWS_BUCKET=BUCKET_NAME
$ export AWS_REGION=REGION_NAME
$ export AWS_ACCESS_KEY_ID=YOUR_AKID
$ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
```

Sample Usage
------------
```shell
$ go run main.go
found a file to upload: d8543ee7-b261-4a57-a572-5541944ff27a.tar.gz
file uploaded to https://test-insights-ibutsu.s3.eu-central-1.amazonaws.com/d8543ee7-b261-4a57-a572-5541944ff27a.tar.gz
```


How to build a binary
-------------------------------
```shell
# for Linux
$ env GOOS=linux go build

# for the current platform
$ go build -o upload-artifact-s3-local
```


