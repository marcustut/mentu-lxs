<script setup lang="ts">
import { useElementVisibility, useTransition } from '@vueuse/core';
import { Dialog, DialogOverlay, TransitionRoot, TransitionChild } from '@headlessui/vue';
import JSConfetti from 'js-confetti';

const score = ref(0);
const main = ref(null);
const pageOne = ref(null);
const pageTwo = ref(null);
const pageThree = ref(null);

const pageTwoAnimationEnd = ref(false);
const pageTwoDialogOpen = ref(false);

const isPageOneVisible = useElementVisibility(pageOne, { scrollTarget: main });
const isPageTwoVisible = useElementVisibility(pageTwo, { scrollTarget: main });
const isPageThreeVisible = useElementVisibility(pageThree, { scrollTarget: main });

const scoreTransition = useTransition(score, {
  duration: 2500,
  transition: [0.75, 0, 0.25, 1],
  onStarted: () => (pageTwoAnimationEnd.value = false),
  onFinished: () => (pageTwoAnimationEnd.value = true),
});

const jsConfetti = new JSConfetti();

// Play Confetti when first page is visible
watch(isPageOneVisible, () => {
  if (isPageOneVisible.value) jsConfetti.addConfetti();
});

// Animate page two on visibility change
watch(isPageTwoVisible, () => {
  if (isPageTwoVisible.value) score.value = 100;
  else score.value = 0;
});
</script>

<template>
  <div ref="main" class="scroll-snap-container fullscreen" font="bahnschrift">
    <div class="item" p="x-4 y-16" bg="white">
      <div
        h="full"
        w="full"
        bg="black"
        pos="relative"
        flex="~ col"
        align="items-center"
        border="rounded-[3rem]"
        justify="center"
        overflow="hidden"
      >
        <img
          src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/CYC%20PK%20Region.svg?alt=media&token=b94ec5d6-151c-4741-9245-aa358ff49864"
          pos="absolute top-4"
          w="14"
          h="14"
        />
        <img
          src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/PageOneBG.svg?alt=media&token=7d17feec-6de0-4c26-9f5f-647b5c198185"
          pos="absolute bottom-0"
          w="1000px"
          h="650px"
          object="cover"
        />
        <div ref="pageOne" m="b-12">
          <p text="4xl" font="bold">CONGRATULATION</p>
          <p m="t-2" text="4xl transparent center" font="jlinxin" class="outline-font">
            门徒练习生
          </p>
        </div>
        <div flex="~" justify="center" pos="relative" w="44" h="44">
          <img
            src="https://i.ibb.co/94JTJrb/dorothea.jpg"
            w="full"
            h="full"
            object="cover"
            border="[#B0F55F] 4 rounded-full"
            shadow="md"
          />
          <div
            p="x-3 y-1"
            pos="absolute -bottom-2"
            bg="white"
            text="xs [#009245]"
            font="bold tracking-wide"
            shadow="lg"
            border="rounded-lg"
          >
            TRAINEE
          </div>
        </div>
        <div m="t-12" z="10" flex="~ col" align="items-center" text="black">
          <p text="2xl" font="bold tracking-wider">DOROTHEA WONG</p>
          <p m="t-1" text="lg" font="tracking-widest">D217427</p>
        </div>
      </div>
    </div>
    <div class="item" p="x-4 y-16" bg="white">
      <div
        p="8"
        w="full"
        h="full"
        flex="~ col"
        text="black"
        align="items-center"
        border="rounded-[3rem]"
        justify="around"
        overflow="hidden"
        gradient="to-br from-[#aff55f] to-[#00ff00]"
      >
        <div ref="pageTwo" flex="~ col" align="items-center">
          <h1 text="5xl" font="tracking-wider jlinxin">成绩单</h1>
          <p text="xs" font="tracking-widest">REPORT CARD</p>
        </div>
        <div w="full">
          <p p="x-4" m="b-1" text="sm" font="bold">FINAL RESULT</p>
          <div
            p="x-4 y-3"
            h="50px"
            bg="white"
            flex="~"
            font="bold"
            border="rounded-full"
            justify="end"
            overflow="hidden"
            :style="{ width: `${scoreTransition}%` }"
          >
            {{ `${scoreTransition.toFixed(0)}%` }}
          </div>
        </div>
        <div w="full">
          <p p="x-4" m="b-1" text="sm" font="bold">COMMENT</p>
          <div
            p="4"
            w="full"
            min-h="364px"
            bg="white"
            flex="~ col"
            font="bold"
            border="rounded-3xl"
            justify="start"
          >
            <TransitionRoot
              :show="pageTwoAnimationEnd"
              enter="transform transition duration-400 ease-in"
              enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100"
              leave="transform duration-200 transition ease-out"
              leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95"
            >
              <TransitionChild
                p="4"
                bg="neonGreen opacity-30"
                border="rounded-3xl"
                enter="transform transition duration-600 ease-in"
                enter-from="opacity-0 translate-y-6"
                enter-to="opacity-100 translate-y-0"
                leave="transform duration-600 transition ease-out"
                leave-from="opacity-100"
                leave-to="opacity-0"
              >
                <p font="bahnschrift">STRENGTH:</p>
                <p font="normal" text="sm">
                  冷静派，管理能力很好，文字也很强，单纯的只相信上帝的计划和安排，即使超乎自己的能力但却能一一完成。
                </p>
              </TransitionChild>
              <TransitionChild
                m="t-4"
                p="4"
                bg="neonGreen opacity-30"
                border="rounded-3xl"
                enter="transform transition duration-600 delay-200 ease-in-out"
                enter-from="opacity-0 translate-y-6"
                enter-to="opacity-100 translate-y-0"
                leave="transform duration-600 transition ease-out"
                leave-from="opacity-100"
                leave-to="opacity-0"
              >
                <p font="bahnschrift">WORDS OF ENCOURAGEMENT:</p>
                <p font="normal" text="sm">
                  更多的操练神给你的能力，不要限制你自己，有时候不一定在你的能力范围内，也不是你能想象的，但我相信神要带领你去到另一个层次。
                </p>
              </TransitionChild>
            </TransitionRoot>
          </div>
        </div>
      </div>
    </div>
    <div class="item" p="x-4 y-16" bg="white">
      <div
        h="full"
        w="full"
        flex="~ col"
        align="items-center"
        border="rounded-[3rem]"
        justify="center"
        overflow="hidden"
        gradient="to-br from-[#333333] to-black"
      >
        <div ref="pageThree"></div>
        <TransitionRoot
          :show="isPageThreeVisible"
          as="div"
          enter="transform transition duration-400 delay-300 ease-in"
          enter-from="opacity-0 scale-80"
          enter-to="opacity-100 scale-100"
          leave="transform duration-200 transition ease-out"
          leave-from="opacity-100 scale-100"
          leave-to="opacity-0 scale-95"
          flex="~ col"
          align="items-center"
        >
          <h3 font="medium">ORGANISED BY</h3>
          <TransitionChild
            enter="transform transition duration-400 delay-[600ms] ease-in"
            enter-from="opacity-0 scale-80 translate-y-6"
            enter-to="opacity-100 scale-100 translate-y-0"
            leave="transform duration-200 transition ease-out"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <img
              src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/CYC%20PK%20Region.svg?alt=media&token=b94ec5d6-151c-4741-9245-aa358ff49864"
              w="36"
              h="36"
            />
          </TransitionChild>
          <TransitionChild
            enter="transform transition duration-400 delay-[800ms] ease-in"
            enter-from="opacity-0 scale-80 translate-y-6"
            enter-to="opacity-100 scale-100 translate-y-0"
            leave="transform duration-200 transition ease-out"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <img
              src="https://firebasestorage.googleapis.com/v0/b/mentu-lxs.appspot.com/o/Logo.png?alt=media&token=a05b3d4c-727e-4f1f-aabe-ae5d2312a1d8"
              w="36"
              h="36"
            />
          </TransitionChild>
        </TransitionRoot>
      </div>
    </div>
  </div>

  <Dialog :open="isPageTwoVisible && pageTwoAnimationEnd && pageTwoDialogOpen">
    <DialogOverlay pos="fixed inset-0" backdrop="~ blur-[2px]" />

    <div
      w="<sm:4/5"
      p="x-6 y-4"
      bg="gray-50"
      pos="fixed top-[50%] left-[50%]"
      flex="~ col"
      text="gray-700"
      align="items-center"
      border="rounded-3xl"
      shadow="2xl"
      transform="~ translate-x-[-50%] translate-y-[-50%]"
      transition="duration-200 ease-in-out"
      class="bg-confetti-animated"
    >
      <p font="bold">Hey, Dorothea Wong!</p>
      <p text="sm">Congratulations for scoring {{ score }}%!</p>

      <button
        m="t-2"
        p="y-2 x-4"
        bg="neonGreen"
        text="xs gray-700"
        font="medium"
        flex="~"
        align="items-center"
        border="rounded-2xl"
        outline="focus:none"
        @click="() => (pageTwoDialogOpen = false)"
      >
        Thanks
        <twemoji:zany-face m="l-1" />
      </button>
    </div>
  </Dialog>
</template>

<style>
.scroll-snap-container {
  display: block;
  overflow-y: scroll;
  overflow-x: hidden;
  -webkit-overflow-scrolling: touch;
  scroll-snap-points-y: repeat(100%);
  scroll-snap-destination: 0 0;
  scroll-snap-type: y mandatory;
  scroll-snap-type: mandatory;
  scroll-behavior: smooth;
}
.scroll-snap-container.horizontal {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  overflow-y: hidden;
  overflow-x: scroll;
  scroll-snap-points-x: repeat(100%);
  scroll-snap-type: x mandatory;
}
.scroll-snap-container.fullscreen {
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  align-items: stretch;
  justify-content: flex-start;
  position: absolute;
  top: 0px;
  bottom: 0px;
  left: 0px;
  right: 0px;
  min-width: 100%;
  min-height: 100%;
}
.scroll-snap-container.fullscreen.horizontal {
  flex-direction: row;
}
.item {
  scroll-snap-align: start;
}
.scroll-snap-container.fullscreen > .item {
  min-height: 100%;
  flex: 1;
}
.scroll-snap-container.horizontal > .item {
  scroll-snap-align: center;
}
.scroll-snap-container.fullscreen.horizontal > .item {
  scroll-snap-align: center;
  min-width: 100%;
  flex: 1;
}
</style>
