<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { useI18n } from 'vue-i18n';
import { firebase, db } from '~/modules/firebase';
import { Field } from 'vee-validate';
import Spinner from '~/components/Spinner.vue';
import type { User } from '~/types';
import type { Ref } from 'vue';

const usersRef = db.collection('users');

const { isAuthenticated, user: firebaseUser } = useAuth(firebase.auth());

const user = ref<User | undefined>();
// @ts-ignore
const users = useFirestore<User[]>(usersRef);

watch(firebaseUser, () => {
  if (!firebaseUser.value) return;

  const userRef = usersRef.where('uid', '==', firebaseUser.value.uid);
  userRef.get().then((query) => {
    if (query.size === 0 || query.size > 1) {
      alert('Something wrong happened, please contact support');
      return;
    }
    query.forEach((doc) => {
      user.value = doc.data() as User;
    });
  });
});

const { locale } = useI18n() as unknown as { locale: Ref<'en' | 'zh-CN'> };
const { t } = useI18n();
</script>

<template>
  <template v-if="isAuthenticated && user">
    <div flex="~" align="items-center">
      <h3 font="bold" text="3xl">{{ t('inbox.title') }}</h3>
      <twemoji:outbox-tray w="7" h="7" m="l-3" />
    </div>

    <form m="t-4" w="full" flex="~ col" align="items-center">
      <div
        p="x-3 y-2"
        bg="gray-200 dark:dark-300"
        text="sm"
        border="rounded-2xl"
        flex="~"
        align="items-center"
      >
        <p bg="neonGreen" border="rounded-2xl" p="x-2 y-1" text="textDark xs" font="medium">
          Sender
        </p>
        <input
          name="sender"
          :placeholder="`${
            user.roles.includes('trainee') ? user.trainee_id : user.roles[0].toUpperCase()
          } - ${locale === 'en' ? user.name.en : user.name['zh-CN']}`"
          m="l-2"
          w="full"
          bg="gray-200 dark:dark-300"
          text="sm placeholder-gray-700 dark:placeholder-gray-200"
          font="bold"
          outline="focus:none"
          disabled
        />
      </div>
      <div
        w="full"
        m="t-4"
        p="x-3 y-2"
        bg="gray-200 dark:dark-300"
        text="sm"
        border="rounded-2xl"
        flex="~"
        align="items-center"
      >
        <p bg="neonGreen" border="rounded-2xl" p="x-2 y-1" text="textDark xs" font="medium">
          Receiver
        </p>
        <Field
          v-if="users"
          as="select"
          name="receiver"
          :placeholder="'WM - Marcus Lee'"
          m="l-2"
          w="full"
          bg="gray-200 dark:dark-300"
          text="sm gray-700 dark:gray-200 placeholder-gray-700 dark:placeholder-gray-200"
          font="bold"
          outline="focus:none"
        >
          <option value="everyone">Everyone</option>
          <option v-for="user in users" :key="user.uid" :value="user.name.en">
            {{ user.name.en }}
          </option>
        </Field>
        <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
      </div>
      <!-- <span m="t-2" text="xs red-500">{{ nameError }}</span> -->
    </form>
  </template>
  <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
</template>

<route lang="yaml">
meta:
  layout: app
</route>
