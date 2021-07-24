<script setup lang="ts">
// import { useRouter } from 'vue-router'
// import { useI18n } from 'vue-i18n'
import { useAuth } from '@vueuse/firebase'
import { firebase, signOut } from '~/modules/firebase'

const { isAuthenticated, user } = useAuth(firebase.auth)

// const router = useRouter()
// const go = () => {
//   if (name.value)
//     router.push(`/hi/${encodeURIComponent(name.value)}`)
// }

// const { t } = useI18n()
</script>

<template>
  <div class="mx-auto">
    <button @click="signOut">
      Sign Out
    </button>
    <pre>isAuthenticated: {{ isAuthenticated }}</pre>
    <pre>{{ user }}</pre>
  </div>
</template>
