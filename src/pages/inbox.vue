<script setup lang="ts">
import { useAuth } from '@vueuse/firebase';
import { storage, firebase } from '~/modules/firebase';
import { useI18n } from 'vue-i18n';
import { useTimeoutFn } from '@vueuse/shared';
import Spinner from '~/components/Spinner.vue';

type File = {
  contentType: string;
  url: string;
};

const { user } = useAuth(firebase.auth());

const urls = ref<File[]>([]);

const { start } = useTimeoutFn(() => {
  getMediaURLs(user.value!.uid);
}, 3000);

watch([user, urls], async () => {
  if (!user.value) return;

  start();
});

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

const { t } = useI18n();
</script>

<template>
  <div flex="~" align="items-center">
    <h3 font="bold" text="3xl">{{ t('inbox.title') }}</h3>
    <twemoji:inbox-tray w="7" h="7" m="l-3" />
  </div>

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
</template>

<route lang="yaml">
meta:
  layout: app
</route>
