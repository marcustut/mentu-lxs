import { useDark, useToggle } from '@vueuse/core';

export const isDark = useDark();
export const toggleDark = useToggle(isDark) as (payload?: MouseEvent) => void;
