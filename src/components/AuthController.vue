<script setup lang="ts">
import { Dialog, DialogOverlay } from '@headlessui/vue'
import { useAuth } from '@vueuse/firebase'
import { useI18n } from 'vue-i18n'
import { useTimeout } from '@vueuse/shared'
import { firebase, signIn } from '~/modules/firebase'
import ThemeToggle from '~/components/ThemeToggle.vue'
import LocaleToggle from '~/components/LocaleToggle.vue'

const { isAuthenticated } = useAuth(firebase.auth())

// Timeout to wait for 'isAuthenticated'
const ready = useTimeout(2000)

const { t } = useI18n()
</script>

<template>
  <Dialog :open="ready && !isAuthenticated">
    <DialogOverlay pos="fixed inset-0" backdrop="~ blur-[2px]" />

    <div
      w="<sm:4/5"
      p="x-6 y-4"
      bg="gray-50 dark:dark-400"
      pos="fixed top-[50%] left-[50%]"
      flex="~ col"
      align="items-center"
      text="gray-700 dark:gray-200"
      border="rounded-3xl"
      shadow="2xl"
      transform="~ translate-x-[-50%] translate-y-[-50%]"
      transition="duration-200 ease-in-out"
    >
      <h4 text="2xl" font="bold">
        {{ t('login.header') }}
      </h4>
      <p m="t-2" text="sm center" font="leading-tight">
        {{ t('login.description') }}
      </p>
      <button
        display="<sm:relative"
        w="<sm:full"
        p="x-6 y-3"
        m="t-4"
        bg="white dark:true-gray-700"
        flex="~"
        align="items-center"
        text="gray-700 dark:gray-50"
        border="rounded-full"
        shadow="md"
        @click="signIn"
      >
        <grommet-icons:google display="<sm:absolute" w="6" h="6" />
        <span w="<sm:full" m="l-2" font="medium">{{ t('login.google_login') }}</span>
      </button>
      <div m="t-4" flex="~" align="items-center">
        <ThemeToggle />
        <LocaleToggle m="l-2" />
      </div>
    </div>
  </Dialog>
</template>
