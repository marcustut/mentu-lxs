<script setup lang="ts">
import { useFirestore } from '@vueuse/firebase';
import { useGroupVideos } from '~/stores';
import { db } from '~/modules/firebase';
import type { PropType } from 'vue';
import type { Group } from '~/types';

const props = defineProps({
  groupId: { type: String, required: true },
  groupName: { type: String, required: true },
  videoName: { type: String, required: true },
  caption: { type: String, default: '' },
  videoSrc: { type: String, required: true },
  postedAt: { type: Date, required: true },
  likes: { type: Number, required: true },
  tags: { type: Array as PropType<string[]>, default: [] },
});

const ONE_DAY = 60 * 60 * 24 * 1000; // in ms

const groupRef = db.collection('groups').doc(props.groupId);

// @ts-ignore
const group = useFirestore<Group>(groupRef);

const groupVideosState = useGroupVideos();

const liked = ref<boolean | undefined>();

const found = Object.keys(groupVideosState.value.likesHistory).find((k) => k === props.groupId);

if (found) {
  liked.value =
    new Date().getTime() - new Date(groupVideosState.value.likesHistory[found]).getTime() < ONE_DAY;
}

const formatDate = (date: Date): string => {
  return `${new Intl.DateTimeFormat('en', {
    hour: '2-digit',
    minute: '2-digit',
    hour12: true,
  }).format(date)} · ${new Intl.DateTimeFormat('en', { month: 'short' }).format(
    date
  )} ${new Intl.DateTimeFormat('en', { day: 'numeric' }).format(date)} ${new Intl.DateTimeFormat(
    'en',
    {
      year: 'numeric',
    }
  ).format(date)}`;
};

const addLike = (groupId: string) => {
  if (!group.value) return;

  groupRef
    .update({
      video: {
        ...group.value.video,
        likes: group.value.video.likes + 1,
      },
    })
    .then(() => {
      groupVideosState.value.likesHistory[groupId] = new Date().toString();
      liked.value = true;
      alert(`Thanks for your vote! (You've voted for ${props.groupName})`);
      return;
    })
    .catch((err) => {
      console.error(err);
      alert(`An error occured! Look for help from support`);
    });
};

const likeHandler = () => {
  if (!group.value) return;

  // If never liked before
  if (!groupVideosState.value.likesHistory[group.value.id]) {
    addLike(group.value.id);
    return;
  }

  const lastLiked = new Date(groupVideosState.value.likesHistory[group.value.id]);

  // If unable to convert to a Date
  if (lastLiked.toString() === 'Invalid Date') {
    console.error('Invalid date string at ' + group.value.id);
    alert(`An error occured! Look for help from support`);
    return;
  }

  // If already liked today
  if (new Date().getTime() - lastLiked.getTime() < ONE_DAY) {
    alert(
      `You have voted for this group today, your next voting time is [${formatDate(
        new Date(lastLiked.getTime() + ONE_DAY)
      )}]`
    );
    return;
  }

  addLike(group.value.id);
};
</script>

<template>
  <div v-if="group" p="3" pos="relative" border="1 gray-300 dark:true-gray-600 rounded-lg">
    <p font="bold">{{ props.groupName }}</p>
    <img
      w="8"
      h="8"
      pos="absolute top-3 right-3"
      object="cover"
      src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/Icon.png?alt=media&token=d72bc58e-b45a-401d-b370-9b961213352c"
    />
    <p text="xs">@{{ props.videoName }}</p>
    <p m="t-2">{{ props.caption }}</p>
    <div flex="~" text="sm neonGreenDark dark:neonGreen" font="medium">
      <span v-for="tag in tags" :key="tag" mr="1" underline="hover:~" cursor="pointer"
        >#{{ tag }}</span
      >
    </div>
    <video m="t-4" w="full" border="rounded-lg" controls>
      <source :src="props.videoSrc" />
    </video>
    <p m="t-2" text="xs">{{ formatDate(props.postedAt) }}</p>
    <div m="y-2" h="[1px]" w="full" bg="gray-300 dark:true-gray-600" />
    <div p="x-1" flex="~" align="items-center">
      <button
        m="r-4"
        flex="~"
        align="items-center"
        :text="liked ? 'red-500' : 'hover:red-500'"
        outline="focus:none"
        class="group"
        @click="likeHandler"
      >
        <div p="1" :bg="!liked && 'group-hover:red-200'" border="rounded-full">
          <ion:heart-outline v-if="!liked" w="5" h="5" />
          <ion:heart-sharp v-else w="5" h="5" />
        </div>
        <span m="l-1 b-1" text="sm" font="medium">{{ props.likes }}</span>
      </button>

      <button
        flex="~"
        align="items-center"
        text="hover:green-500"
        outline="focus:none"
        class="group"
      >
        <div p="1" bg="group-hover:green-200" border="rounded-full">
          <octicon:share-16 w="5" h="5" />
        </div>
        <span m="l-1 b-1" text="sm" font="medium">Share</span>
      </button>
    </div>
  </div>
  <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
</template>
