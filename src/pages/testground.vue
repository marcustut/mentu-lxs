<script setup lang="ts">
import { ref } from 'vue';
import { useElementVisibility } from '@vueuse/core';
const el = ref(null);
const isVisible = useElementVisibility(el);

watch(isVisible, () => {
  console.log(isVisible.value);
});
</script>

<template>
  <div ref="el" style="height: 200px">
    <note class="mb-2"> Info on the right bottom corner </note>
    <div class="area">Target Element (scroll down)</div>
  </div>
  <div class="float">
    {{ isVisible }}
  </div>
  <div h="1000px">heyy</div>
</template>
