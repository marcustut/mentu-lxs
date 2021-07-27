import { createGlobalState, useStorage } from '@vueuse/core'

export const useLaunching = createGlobalState(() => useStorage('mentu-lxs-launching', {
  launched: false,
}))
