<script setup lang="ts">
import { useAuth } from '@vueuse/firebase';
import { firebase } from '~/modules/firebase';

const { user } = useAuth(firebase.auth());

const videos = [
  {
    id: 1,
    title: '总制作人 PD Video',
    video_link: 'https://www.youtube.com/embed/dxbzFSIoqCw',
    date_published: new Date('2021-08-01T20:30:00+08:00'),
    link: '',
  },
  {
    id: 2,
    title: '《有祢》MV',
    video_link: 'https://www.youtube.com/embed/d835iRbjPKA',
    date_published: new Date('2021-07-29T20:30:00+08:00'),
    link: 'https://mentu-lxs.notion.site/MV-42e736076a0144e3805d977f4c3ff5e0',
  },
];
</script>

<template>
  <div
    v-for="video in videos.sort((a, b) => b.id - a.id)"
    m="t-4"
    p="4"
    bg="gray-200 dark:dark-400"
    border="rounded-3xl"
  >
    <h4 text="2xl" font="bold">{{ video.title }}</h4>
    <div m="y-2" w="full" p="t-[56.25%]" pos="relative" overflow="hidden">
      <iframe
        :src="video.video_link"
        title="YouTube video player"
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen
        w="full"
        h="full"
        pos="absolute inset-0"
      ></iframe>
    </div>
    <p m="b-2">发布日期: {{ video.date_published.toLocaleDateString() }}</p>
    <a
      :href="video.link"
      target="_blank"
      flex="~"
      align="items-center"
      class="text-xs text-gray-700 underline dark:text-neonGreen"
      >Notion 帖子 <twemoji-link m="l-2"
    /></a>
  </div>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
