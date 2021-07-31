<script setup lang="ts">
import { Dialog, DialogOverlay } from '@headlessui/vue';
import { useAuth, useFirestore } from '@vueuse/firebase';
import { firebase, db } from '~/modules/firebase';
import type { User } from '~/types';

const usersRef = db.collection('users');

// @ts-ignore
const users = useFirestore<User[]>(usersRef);

const { isAuthenticated, user: firebaseUser } = useAuth(firebase.auth());

const user = ref<User | undefined>();

watch(users, () => {
  if (!users.value) return;
  if (!firebaseUser.value) return;

  user.value = users.value.find((u) => u.uid === firebaseUser.value?.uid);
});
</script>

<template>
  <Dialog :open="isAuthenticated && !user">
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
      Sign Up
      <button>Hi</button>
    </div>
  </Dialog>
</template>
