import { useFlag, useFlagsStatus } from '@unleash/proxy-client-react';
import useChrome from '@redhat-cloud-services/frontend-components/useChrome';

const useFeatureFlag = (flag) => {
  const { flagsReady } = useFlagsStatus();
  const isFlagEnabled = useFlag(flag);
  return flagsReady ? isFlagEnabled : false;
};

export default useFeatureFlag;

export const UPDATE_RISKS_ENABLE_FLAG =
  'ocp-advisor.upgrade-risks.enable-in-stable';

export const useUpdateRisksFeatureFlag = () => {
  const updateRisksEnabled = useFeatureFlag(UPDATE_RISKS_ENABLE_FLAG);
  const chrome = useChrome();

  return chrome.isBeta() || updateRisksEnabled;
};
