<script setup lang="ts">
const links: {
  img: string;
  title: string;
  link: string;
  disabled?: boolean;
  onclick?: () => void;
}[] = [
  {
    img: 'https://i.pinimg.com/favicons/f446bd086a965df9beea6b3962eaa174079c624972e899cf3a39331e.png?d23e369be383e6616adf07b90a351957',
    title: '追梦的休息站 Zoom',
    link: 'https://zoom.us/j/6905199217?pwd=bmdkTkVSUlZuNjNLTnFMVW5RVnQzQT09',
  },
  {
    img: 'https://i.pinimg.com/favicons/f446bd086a965df9beea6b3962eaa174079c624972e899cf3a39331e.png?d23e369be383e6616adf07b90a351957',
    title: '你的装备已更新了吗？Zoom',
    link: 'https://zoom.us/j/6905199217?pwd=bmdkTkVSUlZuNjNLTnFMVW5RVnQzQT09',
  },
  {
    img: 'https://i.pinimg.com/favicons/f446bd086a965df9beea6b3962eaa174079c624972e899cf3a39331e.png?d23e369be383e6616adf07b90a351957',
    title: '你是属于什么担当？Zoom',
    link: 'https://zoom.us/j/6905199217?pwd=bmdkTkVSUlZuNjNLTnFMVW5RVnQzQT09',
  },
  {
    img: 'https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/MVThumbnail.PNG?alt=media&token=0dae8fc6-ad79-4eee-a272-33de46f52d5a',
    title: '《有祢》专属 MV',
    link: 'https://mentu-lxs.notion.site/MV-42e736076a0144e3805d977f4c3ff5e0',
  },
  {
    img: 'data:image/png;base64,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',
    title: 'Notion',
    link: 'https://mentu-lxs.notion.site/e45ced448d8b485d87063d901af7917b',
  },
];
</script>

<template>
  <div w="screen" h="screen" bg="textDark">
    <div
      overflow="y-hidden"
      display="flex"
      flex="col"
      justify="center"
      align="items-center"
      p="x-4"
      text="light-50"
      w="max-2xl"
      m="x-auto"
      h="full"
    >
      <img
        src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/Icon.png?alt=media&token=d72bc58e-b45a-401d-b370-9b961213352c"
        w="20"
        h="20"
        object="cover"
      />
      <h3 font="bold" m="t-4">门徒练习生</h3>
      <a
        v-for="link in links"
        :key="link.title"
        :href="!link.disabled ? link.link : undefined"
        target="_blank"
        class="text-textDark hover:text-white"
        transition="duration-200 ease-in-out"
        bg="white hover:textDark"
        border="2 white dashed"
        w="full"
        m="t-6"
        p="x-4 y-6"
        display="flex"
        justify="center"
        align="items-center"
        pos="relative"
        cursor="pointer"
        :onclick="link.onclick ? link.onclick : undefined"
      >
        <img :src="link.img" pos="absolute left-4" w="12" h="12" object="cover" />
        <p m="l-4" font="medium">
          {{ link.title }}
        </p>
      </a>
    </div>
  </div>
</template>
