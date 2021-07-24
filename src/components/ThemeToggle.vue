<script setup lang="ts">
import { isDark, toggleDark } from '~/logic/dark'

const props = defineProps({
  class: String,
})
</script>

<template>
  <button :class="props.class" outline="focus:outline-none" @click="toggleDark()">
    <bx-bxs-moon v-if="isDark" />
    <fa-solid:sun v-else />
  </button>
</template>
