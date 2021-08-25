import { createGlobalState, useStorage } from '@vueuse/core';

export const useLaunching = createGlobalState(() =>
  useStorage('mentu-lxs-launching', {
    voted: false,
  })
);

type GroupVideosState = {
  likesHistory: Record<string, string>;
};

export const useGroupVideos = createGlobalState(() =>
  useStorage<GroupVideosState>('mentu-lxs-videos', { likesHistory: {} })
);
