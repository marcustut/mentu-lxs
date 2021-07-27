<script setup lang="ts">
import AuthController from '~/components/AuthController.vue'
import Navbar from '~/components/Navbar.vue'
</script>

<template>
  <main text="gray-700 dark:gray-200">
    <AuthController />
    <Navbar class="p-4" />
    <div class="px-4">
      <router-view />
    </div>
  </main>
</template>
