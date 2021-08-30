<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { useTimeoutFn } from '@vueuse/shared';
import { useAuth, useFirestore } from '@vueuse/firebase';
import { Dialog, DialogOverlay } from '@headlessui/vue';
import { storage, firebase, db } from '~/modules/firebase';
import Spinner from '~/components/Spinner.vue';
import type { User, InboxMessage } from '~/types';

type File = {
  contentType: string;
  url: string;
};

type InboxMessageWithUser = {
  attachment: string[];
  message: string;
  receiver: User | null;
  sender: User | null;
  created_at: firebase.firestore.Timestamp;
};

const messagesRef = db.collection('messages');

const { user } = useAuth(firebase.auth());
// @ts-ignore
const messages = useFirestore<InboxMessage[]>(messagesRef);
const messagesWithUsers = ref<InboxMessageWithUser[]>();
const currentMessage = ref<InboxMessageWithUser>();

const urls = ref<File[]>([]);

const { start } = useTimeoutFn(() => {
  getMediaURLs(user.value!.uid);
}, 3000);

watch([user, urls], async () => {
  if (!user.value) return;
  start();
});

watch(messages, async () => {
  if (!messages.value) return;
  if (!user.value) return;

  messagesWithUsers.value = messages.value.filter((m) => m.receiver === user.value?.uid);

  for (let i = 0; i < messagesWithUsers.value.length; i++) {
    const { data: receiver } = await getUser(
      messagesWithUsers.value[i].receiver as unknown as string
    );
    const { data: sender } = await getUser(messagesWithUsers.value[i].sender as unknown as string);
    messagesWithUsers.value[i] = {
      ...messagesWithUsers.value[i],
      receiver: receiver,
      sender,
    };
  }
});

const getUser = async (uid: string) => {
  const querySnapshot = await db.collection('users').where('uid', '==', uid).get();
  let user = null;

  querySnapshot.forEach((doc) => {
    if (doc.exists) user = doc.data();
  });

  if (!user) return { data: null, error: new Error('Unable to fetch') };

  return { data: user, error: null };
};

const getMediaURLs = async (uid: string) => {
  const res = await storage.ref('PrayerVideo').listAll();

  // Find the person's folder from the list
  res.prefixes.forEach(async (ref) => {
    if (ref.name.includes(uid)) {
      const folders = await ref.listAll();
      folders.items.forEach(async (item) => {
        const metadata = await item.getMetadata();
        if (metadata) {
          if (
            (metadata.contentType as string).includes('video') ||
            (metadata.contentType as string).includes('audio')
          ) {
            const url = await item.getDownloadURL();
            urls.value.push({ url: url, contentType: metadata.contentType } as File);
          }
        }
      });
    }
  });
};

const unsetCurrentMessage = () => {
  currentMessage.value = undefined;
};

const { t } = useI18n();
</script>

<template>
  <div flex="~" align="items-center">
    <h3 font="bold" text="3xl">{{ t('inbox.title') }}</h3>
    <twemoji:inbox-tray w="7" h="7" m="l-3" />
  </div>

  <div v-if="messagesWithUsers && messages" m="t-4">
    <button
      v-for="(message, index) in messagesWithUsers"
      :m="index !== 0 && 't-2'"
      p="x-3 y-2"
      bg="gray-200 dark:dark-300"
      text="sm"
      flex="~"
      border="rounded-2xl"
      align="items-center"
      outline="focus:none"
      ring="focus:2"
      @click="
        () => {
          currentMessage = message;
        }
      "
    >
      <template v-if="message.sender">
        <img
          border="rounded-full"
          object="cover"
          w="8"
          h="8"
          :src="
            message.sender.avatar_url
              ? message.sender.avatar_url
              : `https://ui-avatars.com/api/?name=${message.sender.name.en}`
          "
        />
        <div w="full" m="l-3" flex="~ col" align="items-start">
          <p w="full" font="medium" text="lg left" class="line-clamp-1">{{ message.message }}</p>
          <p class="line-clamp-1">
            by <span font="medium">{{ message.sender.name.en }}</span>
          </p>
        </div>
      </template>
      <mdi:chevron-right m="l-auto" w="8" h="8" text="dark-300 dark:gray-200" />
    </button>
  </div>
  <Spinner v-else mx="auto" my="12" animate="spin" w="12" h="12" text="neonGreen" />

  <div v-if="urls.length !== 0" m="t-8" flex="~ col" justify="center" align="center">
    <div v-for="(url, index) in urls" :key="url.url" :m="index !== 0 ? 't-4' : ''">
      <video v-if="url.contentType.includes('video')" w="300px" controls>
        <source :src="url.url" :type="url.contentType" />
      </video>
      <audio v-else-if="url.contentType.includes('audio')" controls>
        <source :src="url.url" :type="url.contentType" />
      </audio>
    </div>
  </div>
  <Spinner v-else mx="auto" my="12" animate="spin" w="12" h="12" text="neonGreen" />

  <Dialog :open="currentMessage != undefined" @close="unsetCurrentMessage">
    <DialogOverlay pos="fixed inset-0" backdrop="~ blur-[2px]" />

    <div
      w="4/5"
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
      <div
        w="full"
        p="x-3 y-2"
        bg="gray-200 dark:dark-300"
        text="sm"
        border="rounded-2xl"
        flex="~"
        align="items-center self-start"
      >
        <p bg="neonGreen" border="rounded-2xl" p="x-2 y-1" text="textDark xs" font="medium">From</p>
        <p
          m="l-2"
          w="full"
          bg="gray-200 dark:dark-300"
          text="sm placeholder-gray-700 dark:placeholder-gray-200"
          font="bold"
        >
          {{ currentMessage?.sender?.name.en }}
        </p>
      </div>
      <p m="t-4" p="4" bg="gray-200 dark:dark-300" text="space-pre-wrap" border="rounded-2xl">
        {{ currentMessage?.message.split('\\n').join('\n') }}
      </p>
      <div
        v-if="currentMessage?.attachment.length !== 0"
        w="full"
        m="t-4"
        flex="~ col"
        justify="center"
        align="center"
      >
        <a
          v-for="(url, index) in currentMessage?.attachment"
          :key="url"
          :m="index !== 0 ? 't-4' : ''"
          bg="gray-200 hover:neonGreen dark:dark-300"
          p="x-3 y-2"
          flex="~"
          font="medium"
          align="items-center"
          border="rounded-xl"
          class="text-sm hover:text-textDark"
          :href="url"
          target="_blank"
        >
          <jam:files w="6" h="6" m="r-2" />
          Attachment {{ index }}
        </a>
      </div>
      <button
        w="full"
        m="t-4"
        p="x-3 y-2"
        bg="gray-200 hover:red-500 dark:dark-300"
        text="hover:textDark"
        font="medium"
        flex="~"
        align="items-center"
        border="rounded-3xl"
        justify="center"
        outline="focus:none"
        @click="unsetCurrentMessage"
      >
        <jam:close-circle w="6" h="6" m="r-1" />
        <span m="b-1">Close</span>
      </button>
    </div>
  </Dialog>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
