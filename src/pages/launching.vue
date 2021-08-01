<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { useSound } from '@vueuse/sound';
// @ts-ignore
import vueDanmaku from 'vue3-danmaku';
import { Dialog, DialogOverlay } from '@headlessui/vue';
import { onStartTyping, useTimeoutFn, useNow } from '@vueuse/core';
import { useI18n } from 'vue-i18n';
import { firebase, db } from '~/modules/firebase';
import { useLaunching } from '~/stores';
import Spinner from '~/components/Spinner.vue';
import Firework from '~/components/Firework.vue';
import type { Message, Launching } from '~/types';

const launchingRef = db.collection('interaction').doc('launching');
const messagesRef = db.collection('interaction').doc('launching').collection('messages');

const { user } = useAuth(firebase.auth());
const { play: playSFX } = useSound(
  'https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/FireworkSFX.mp3?alt=media&token=b944b420-7cf4-4ef9-99ed-22e8c293192b',
  { volume: 0.2 }
);
const now = useNow();
const { t } = useI18n();
// @ts-ignore
const launching = useFirestore<Launching>(launchingRef);
// @ts-ignore
const messages = useFirestore<Message[]>(messagesRef);

const messageText = ref('');
const launched = useLaunching();
const dialogOpen = ref(false);
const startFirework = ref(false);
const inputRef = ref<HTMLInputElement | null>(null);
const countdownStr = ref('');

// Watcher only start confetti if the bar first reach 100%
watch(launching, () => {
  if (!launching.value) return;
  if (launching.value.end_datetime.toDate() >= now.value) return;

  const percentage = (launching.value.count / launching.value.participants) * 100;

  if (percentage < 100) {
    startFirework.value = false;
    return;
  }

  startFirework.value = true;
  playSFX();
});

// Update countdown duration
watch([launching, now], () => {
  if (!launching.value) return;

  const diff = launching.value.start_datetime.toDate().getTime() - now.value.getTime();

  const totalNanoseconds = ((3e5 - diff) / 1000) >> 0;
  const totalMinutes = (totalNanoseconds / 60) >> 0;
  const hours = (totalMinutes / 60) >> 0;
  const minutes = ((totalMinutes / 60 - hours) * 60) >> 0;
  const seconds = totalNanoseconds - totalMinutes * 60;

  countdownStr.value = `${Math.abs(hours).toString().padStart(2, '0')}:${Math.abs(minutes)
    .toString()
    .padStart(2, '0')}:${Math.abs(seconds).toString().padStart(2, '0')}`;
});

// Sort messages by latest date
watch(messages, () =>
  messages.value?.sort((a, b) => a.createdAt.toDate().getTime() - b.createdAt.toDate().getTime())
);

const focusInput = useTimeoutFn(() => {
  inputRef.value?.focus();
}, 0);

onStartTyping(() => {
  // @ts-ignore
  if (!inputRef.value?.active) focusInput.start();
});

const openDialog = () => (dialogOpen.value = true);
const closeDialog = () => {
  dialogOpen.value = false;
  focusInput.start();
};

const addCount = () => {
  if (launched.value.voted) {
    openDialog();
    return;
  }

  launchingRef
    .update({
      count: firebase.firestore.FieldValue.increment(1),
    })
    .then(() => {
      openDialog();
      launched.value.voted = true;
    });
};

const sendMessage = () => {
  if (messageText.value !== '') {
    messagesRef.add({
      content: messageText.value,
      sentBy: { name: user.value?.displayName, avatar_url: user.value?.photoURL },
      createdAt: new Date(),
    });
    messageText.value = '';
  }
};
</script>

<template>
  <template v-if="launching">
    <template v-if="launching.end_datetime.toDate() >= now">
      <Firework v-if="startFirework" />

      <div
        v-if="launching.start_datetime.toDate() < now"
        h="[91vh]"
        flex="~ col"
        justify="center"
        align="items-center"
        pos="relative"
      >
        <div v-if="launched.voted" w="full" p="t-[56.25%]" pos="relative" overflow="hidden">
          <iframe
            src="https://www.youtube.com/embed/oJAxa_ODcSc"
            title="YouTube video player"
            frameborder="0"
            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
            allowfullscreen
            w="full"
            h="full"
            pos="absolute inset-0"
          ></iframe>
        </div>

        <div h="2" w="72" :m="launched.voted ? 't-6' : ''" pos="relative">
          <div
            h="full"
            bg="neonGreen"
            border="rounded"
            pos="absolute"
            :style="{
              width: `${
                (launching.count / launching.participants) * 100 >= 100
                  ? '100'
                  : ((launching.count / launching.participants) * 100).toFixed(0)
              }%`,
            }"
          ></div>
          <div h="full" w="full" bg="neonGreenShade" border="rounded"></div>
        </div>

        <p m="t-4" text="4xl" font="mono">
          {{ ((launching.count / launching.participants) * 100).toFixed(0) }}%
        </p>

        <input
          ref="inputRef"
          v-model="messageText"
          :placeholder="t('launching.input_placeholder')"
          autofocus
          m="y-4"
          w="2/3 sm:1/3"
          p="x-4 y-3"
          bg="true-gray-50 dark:true-gray-700"
          text="sm"
          border="rounded-2xl"
          shadow="lg"
          outline="focus:none"
          ring="focus:2 neonGreen opacity-70"
        />
        <button
          p="x-4 y-3"
          bg="neonGreen"
          text="sm dark-700"
          font="medium"
          flex="~"
          align="items-center"
          border="rounded-2xl"
          outline="focus:none"
          shadow="lg"
          @click="sendMessage"
        >
          {{ t('launching.button') }}
          <twemoji-outbox-tray w="4" h="4" m="l-1" />
        </button>

        <div v-if="messages" pos="absolute top-0">
          <vue-danmaku :danmus="messages" use-slot w="screen" h="[250px]" :speeds="100">
            <template #dm="{ danmu }">
              <div
                p="x-3 y-2"
                bg="gray-100 dark:true-gray-700"
                border="rounded-full"
                flex="~"
                align="items-center"
              >
                <img
                  border="rounded-full"
                  w="6"
                  h="6"
                  m="r-1"
                  :src="danmu.sentBy.avatar_url"
                  :alt="danmu.sentBy.name"
                />
                <span text="gray-700 dark:gray-200 sm"
                  >{{ danmu.sentBy.name }}: {{ danmu.content }}</span
                >
              </div>
            </template>
          </vue-danmaku>
        </div>
      </div>

      <div v-else h="[85vh]" flex="~ col" justify="center" align="items-center">
        <div text="sm" flex="~" align="items-center">
          <noto-stopwatch w="4" h="4" m="r-1" />Launching will be available in
        </div>
        <p font="mono bold" text="3xl">{{ countdownStr }}</p>
      </div>

      <button
        p="x-4 y-3"
        z="10"
        pos="fixed bottom-4 right-4"
        bg="neonGreen"
        border="rounded-2xl"
        shadow="xl"
        text="xl"
        outline="focus:none"
        @click="addCount"
      >
        <twemoji-party-popper />
      </button>

      <Dialog :open="launched.voted && dialogOpen" @close="closeDialog">
        <DialogOverlay z="10" pos="fixed inset-0" backdrop="~ blur-[2px]" />

        <div
          w="<sm:4/5"
          p="x-6 y-4"
          z="10"
          bg="gray-50 dark:dark-400"
          pos="fixed top-[50%] left-[50%]"
          flex="~ col"
          align="items-center"
          text="gray-700 dark:gray-200"
          border="rounded-3xl"
          shadow="2xl"
          transform="~ translate-x-[-50%] translate-y-[-50%]"
          transition="duration-200 ease-in-out"
          class="bg-confetti-animated"
        >
          <div flex="~" align="items-center">
            <h4 text="lg center" font="bold" m="r-2">
              {{ t('launching.dialog.title') }}
            </h4>
            <twemoji-zany-face />
          </div>
          <button
            m="t-2"
            p="x-4 y-3"
            bg="gray-200 dark:dark-200"
            text="sm"
            font="medium"
            border="rounded-2xl"
            outline="focus:none"
            @click="closeDialog"
          >
            {{ t('launching.dialog.button') }}
          </button>
        </div>
      </Dialog>
    </template>

    <div>Event is expired</div>
  </template>

  <div v-else h="[91vh]" flex="~ col" justify="center" align="items-center">
    <Spinner animate="spin" w="12" h="12" text="neonGreen" />
  </div>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
