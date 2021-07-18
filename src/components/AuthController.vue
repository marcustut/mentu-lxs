<script setup lang="ts">
import { Dialog, DialogOverlay } from '@headlessui/vue'
import { useAuth } from '@vueuse/firebase'
import { useDark, useToggle } from '@vueuse/core'
import { firebase, signIn } from '~/modules/firebase'

const { isAuthenticated } = useAuth(firebase.auth())

const isDark = useDark()
const toggleTheme = useToggle(isDark)

</script>

<template>
  <Dialog :open="!isAuthenticated">
    <DialogOverlay class="fixed inset-0 backdrop-filter backdrop-blur-[2px]" />

    <div
      class="fixed flex flex-col bg-gray-50 dark:bg-dark-400 top-[50%] left-[50%] transform translate-x-[-50%] translate-y-[-50%] text-gray-700 dark:text-gray-200"
    >
      Sign In
      <button @click="signIn">Sign in with Google</button>
      <button @click="toggleTheme()">Toggle Light/Dark</button>
    </div>
  </Dialog>
</template>
