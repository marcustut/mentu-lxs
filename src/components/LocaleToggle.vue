<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const props = defineProps({
  class: String,
})

const { locale } = useI18n()

const toggleLocale = () => {
  if (locale.value === 'en')
    locale.value = 'zh-CN'
  else
    locale.value = 'en'
}
</script>

<template>
  <button :class="props.class" outline="focus:outline-none" @click="toggleLocale()">
    <uil-english-to-chinese />
  </button>
</template>
