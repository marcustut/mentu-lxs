<script setup lang="ts">
const props = defineProps({
  fullscreen: {
    type: Boolean,
    default: false,
    required: false,
  },
  horizontal: {
    type: Boolean,
    default: false,
    required: false,
  },
});
</script>

<template>
  <div
    class="scroll-snap-container"
    :class="{ fullscreen: props.fullscreen, horizontal: props.horizontal }"
  >
    <slot></slot>
  </div>
</template>
