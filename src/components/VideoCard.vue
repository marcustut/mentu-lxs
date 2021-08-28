<script setup lang="ts">
import { Popover, PopoverButton, PopoverOverlay, PopoverPanel } from '@headlessui/vue';
import { useFirestore } from '@vueuse/firebase';
import { useGroupVideos } from '~/stores';
import { db } from '~/modules/firebase';
import type { PropType } from 'vue';
import type { Group } from '~/types';
import { useClipboard } from '@vueuse/core';

const props = defineProps({
  groupId: { type: String, required: true },
  groupName: { type: String, required: true },
  videoName: { type: String, required: true },
  caption: { type: String, default: '' },
  videoSrc: { type: String, required: true },
  postedAt: { type: Date, required: true },
  likes: { type: Number, required: true },
  tags: { type: Array as PropType<string[]>, default: [] },
  theme: { type: String, required: true },
});

const ONE_DAY = 60 * 60 * 24 * 1000; // in ms

const groupRef = db.collection('groups').doc(props.groupId);

// @ts-ignore
const group = useFirestore<Group>(groupRef);

const groupVideosState = useGroupVideos();

const liked = ref<boolean | undefined>();

const { copy } = useClipboard();

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

const copyHandler = () => {
  copy(encodeURI(`https://mentu-lxs.netlify.app/video/vote#${props.groupId}`));
  alert(`Copied to clipboard`);
};
</script>

<template>
  <div
    :id="group.id"
    v-if="group"
    p="3"
    pos="relative"
    border="1 gray-300 dark:true-gray-600 rounded-lg"
  >
    <p w="4/5" font="bold">{{ props.videoName }}</p>
    <img
      w="8"
      h="8"
      pos="absolute top-3 right-3"
      object="cover"
      src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/Icon.png?alt=media&token=d72bc58e-b45a-401d-b370-9b961213352c"
    />
    <p text="xs">@{{ props.groupName }}</p>
    <div flex="~">
      <p m="t-2" text="xs textDark" font="bold" p="x-2 y-1" bg="neonGreen" border="rounded-2xl">
        主题: {{ props.theme }}
      </p>
    </div>
    <p m="t-2" text="space-pre-wrap">{{ props.caption.split('\\n').join('\n') }}</p>
    <div m="t-1" flex="~ wrap" text="sm neonGreenDark dark:neonGreen" font="medium">
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

      <Popover v-slot="{ open }" m="l-auto" pos="relative">
        <PopoverButton
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
        </PopoverButton>

        <PopoverOverlay class="bg-black" :class="open ? 'opacity-30 fixed inset-0' : 'opacity-0'" />

        <PopoverPanel pos="absolute z-10 -top-28 -left-12">
          <div p="2" bg="green-700" text="textDark" border="rounded-xl">
            <a
              p="2"
              bg="hover:neonGreenDark hover:opacity-25"
              flex="~"
              align="items-center"
              font="bold"
              border="rounded-xl"
              class="text-sm text-white"
              :href="
                encodeURI(
                  `https://api.whatsapp.com/send?text=*Vote for ${group.name} now!*\n\n_https://mentu-lxs.netlify.app/video/vote#${group.id} 🔗_`
                )
              "
              data-action="share/whatsapp/share"
            >
              <logos:whatsapp w="6" h="6" m="r-2" />
              WhatsApp
            </a>
            <button
              p="2"
              bg="hover:neonGreenDark hover:opacity-25"
              flex="~"
              align="items-center"
              font="bold"
              border="rounded-xl"
              text="sm white"
              outline="focus:none"
              @click="copyHandler"
            >
              <mdi:link-variant w="6" h="6" m="r-2" />
              Copy Link
            </button>
          </div>
        </PopoverPanel>
      </Popover>
    </div>
  </div>
  <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
</template>
