<script setup lang="ts">
import { useHead } from '@vueuse/head'

// https://github.com/vueuse/head
// you can use this to manipulate the document head in any components,
// they will be rendered correctly in the html results with vite-ssg
useHead({
  title: '门徒练习生',
  meta: [
    { name: 'description', content: '新一代创造力与影响力' },
  ],
})
</script>

<template>
  <router-view />
</template>
