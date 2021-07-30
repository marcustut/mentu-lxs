<script setup lang="ts">
import { useRouter } from 'vue-router';
import AuthController from '~/components/AuthController.vue';
import Navbar from '~/components/Navbar.vue';

const { currentRoute } = useRouter();
</script>

<template>
  <main text="gray-700 dark:gray-200">
    <AuthController />
    <Navbar :class="`${currentRoute.fullPath === '/' ? 'text-gray-700' : ''} p-4`" />
    <div p="x-4">
      <router-view />
    </div>
  </main>
</template>
