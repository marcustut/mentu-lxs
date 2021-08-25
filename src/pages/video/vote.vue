<script setup lang="ts">
import { useFirestore } from '@vueuse/firebase';
import ThemeToggle from '~/components/ThemeToggle.vue';
import VideoCard from '~/components/VideoCard.vue';
import { db } from '~/modules/firebase';
import type { Group } from '~/types';

const groupsRef = db.collection('groups');

// @ts-ignore
const groups = useFirestore<Group[]>(groupsRef);
</script>

<template>
  <div p="4" text="dark-700 dark:light-200">
    <div flex="~" justify="end">
      <ThemeToggle />
    </div>
    <div m="y-4" flex="~" justify="center" align="items-center">
      <p>Vote for your group of choice below!</p>
    </div>
    <VideoCard
      v-if="groups"
      v-for="(group, index) in groups"
      :key="group.id"
      :group-id="group.id"
      :group-name="group.name"
      :video-name="group.video.name"
      :caption="group.video.caption"
      :posted-at="group.video.postedAt.toDate()"
      :video-src="group.video.src"
      :tags="group.video.tags"
      :likes="group.video.likes"
      :liked="false"
      :m="`x-auto ${index === 0 ? '' : 't-4'}`"
      w="sm:1/2 lg:1/3 full"
    />
    <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
  </div>
</template>
