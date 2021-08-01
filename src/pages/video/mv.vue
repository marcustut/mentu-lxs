<script setup lang="ts">
// import ComingSoon from '~/components/ComingSoon.vue'
</script>

<template>
  <div class="flex h-screen justify-center items-center">
    <div w="full" p="t-[56.25%]" pos="relative" overflow="hidden">
      <iframe
        src="https://www.youtube.com/embed/d835iRbjPKA"
        title="YouTube video player"
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen
        w="full"
        h="full"
        pos="absolute inset-0"
      ></iframe>
    </div>
  </div>
  <!-- <ComingSoon />
  <div
    p="t-32"
    text="xl center neonGreen"
    font="bold italic mono"
    min-h="screen"
    flex="~"
    justify="center"
    align="items-center"
    overflow="y-hidden"
  >
    Coming Soon
  </div> -->
</template>
