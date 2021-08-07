<script setup lang="ts">
import { db } from '~/modules/firebase';
import { useFirestore } from '@vueuse/firebase';
import type { User } from '~/types';

const usersRef = db.collection('users');

// @ts-ignore
const userResponses = useFirestore<User[]>(usersRef);

watch(userResponses, () => {
  userResponses.value?.forEach((u) => console.log(`${u.displayName}: ${u.uid}`));
});
</script>

<template>
  <pre>{{ userResponses }}</pre>
</template>
