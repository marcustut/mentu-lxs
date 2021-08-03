<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { useRouter } from 'vue-router';
import { firebase, db } from '~/modules/firebase';
import { disc as data, personality } from '~/data';
import { useI18n } from 'vue-i18n';
import type { Ref } from 'vue';
import type { Disc } from '~/types';

const discResponsesRef = db.collection('disc-responses');

// @ts-ignore
const discResponses = useFirestore<Disc[]>(discResponsesRef);

const { user } = useAuth(firebase.auth());
const { push } = useRouter();
const { locale } = useI18n() as unknown as { locale: Ref<string> };
const { t } = useI18n();

const disc = ref<Disc | undefined>();

watch(discResponses, () => {
  if (!discResponses.value) return;
  if (!user.value) return;

  disc.value = discResponses.value.find((r) => r.uid === user.value?.uid);
});

const getTrait = (type: 'd' | 'i' | 's' | 'c') => {
  if (!disc.value) return;
  const answers = Object.values(data).map((d) => d.answer);
  const scores = Object.values(disc.value.responses).map((ans, index) => answers[index][ans]);
  return scores.filter((s) => s === type).length;
};

const getPersonality = () => {
  if (!disc.value) return;

  const traits = {
    d: getTrait('d')!,
    i: getTrait('i')!,
    s: getTrait('s')!,
    c: getTrait('c')!,
  };

  const highestTrait = Object.entries(traits).reduce((a, b) => {
    const max = Math.max(a[1], b[1]);
    return a[1] === max ? a : b;
  });

  return personality[highestTrait[0]];
};
</script>

<template>
  <div
    p="4"
    flex="~ col"
    justify="center"
    align="items-center"
    bg="indigo-700"
    border="rounded-3xl"
    class="bg-confetti-static"
  >
    <h3 m="b-2" text="xl" font="bold">{{ t('test.disc.title') }}</h3>
    <template v-if="disc">
      <h4 m="b-2" font="medium">{{ t('test.disc.result') }}</h4>
      <table text="xl">
        <thead>
          <tr bg="blue-700">
            <th p="2" border="~">D</th>
            <th p="2" border="~">I</th>
            <th p="2" border="~">S</th>
            <th p="2" border="~">C</th>
          </tr>
        </thead>
        <tbody>
          <tr bg="blue-600">
            <td p="2" border="~">{{ getTrait('d') }}</td>
            <td p="2" border="~">{{ getTrait('i') }}</td>
            <td p="2" border="~">{{ getTrait('s') }}</td>
            <td p="2" border="~">{{ getTrait('c') }}</td>
          </tr>
        </tbody>
      </table>
      <p m="t-2" text="sm center">
        {{ t('test.disc.description') }}
        <span font="bold">{{
          locale === 'en' ? getPersonality()?.en : getPersonality()?.['zh-CN']
        }}</span>
      </p>
    </template>
    <div v-else-if="discResponses && user && !disc" flex="~ col" align="items-center">
      <p flex="~" align="items-center">
        {{ t('test.disc.no_test_done') }}
        <twemoji-grimacing-face m="l-1.5" w="4" h="4" />
      </p>
      <button
        m="t-4 b-1"
        p="x-4 y-3"
        bg="neonGreen"
        text="gray-700 sm"
        font="medium"
        border="rounded-2xl"
        outline="focus:none"
        @click="push('/test/disc')"
      >
        {{ t('test.disc.button') }}
      </button>
    </div>
    <Spinner v-else animate="spin" m="y-8 x-auto" w="12" h="12" text="neonGreen" />
  </div>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
