<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { useRouter } from 'vue-router';
import { firebase, db } from '~/modules/firebase';
import { giftLabels } from '~/data';
import { useI18n } from 'vue-i18n';
import type { Gift } from '~/types';

const giftResponsesRef = db.collection('gift-responses');

// @ts-ignore
const giftResponses = useFirestore<Gift[]>(giftResponsesRef);

const { user } = useAuth(firebase.auth());
const { push } = useRouter();
const { t } = useI18n();

const gift = ref<Gift | undefined>();

watch(giftResponses, () => {
  if (!giftResponses.value) return;
  if (!user.value) return;

  gift.value = giftResponses.value.find((r) => r.uid === user.value?.uid);
});

const getGift = (type: keyof typeof giftLabels) => {
  if (!gift.value) return;

  let score = 0;

  const locations = giftLabels[type];

  const answers = Object.entries(gift.value.responses).filter((r) =>
    locations.includes(parseInt(r[0]))
  );

  answers.forEach((a) => (score += a[1]));

  return score;
};

const getHighestGiftLabel = () => {
  if (!gift.value) return;

  const scores = {} as Record<keyof typeof giftLabels, number>;

  const gifts = Object.keys(giftLabels) as (keyof typeof giftLabels)[];

  gifts.forEach((g) => (scores[g] = getGift(g)!));

  const arr = Object.entries(scores)
    .filter((s) => s[1] == Math.max(...Object.values(scores)))
    .reverse();

  return arr[0][0];
};
</script>

<template>
  <div
    p="4"
    bg="indigo-700"
    text="gray-200"
    flex="~ col"
    justify="center"
    align="items-center"
    border="rounded-3xl"
    class="bg-confetti-static"
  >
    <h3 m="b-2" text="xl" font="bold">{{ t('test.gift.title') }}</h3>
    <template v-if="gift">
      <h4 m="b-2" font="medium">{{ t('test.gift.result') }}</h4>
      <table text="xl sm">
        <thead>
          <tr bg="blue-700">
            <th p="2" font="normal" border="~">宣教</th>
            <th p="2" font="normal" border="~">传福音</th>
            <th p="2" font="normal" border="~">接待</th>
            <th p="2" font="normal" border="~">信心</th>
            <th p="2" font="normal" border="~">领导</th>
          </tr>
        </thead>
        <tbody>
          <tr bg="blue-600">
            <td p="2" text="center" border="~">{{ getGift('宣教') }}</td>
            <td p="2" text="center" border="~">{{ getGift('传福音') }}</td>
            <td p="2" text="center" border="~">{{ getGift('接待') }}</td>
            <td p="2" text="center" border="~">{{ getGift('信心') }}</td>
            <td p="2" text="center" border="~">{{ getGift('领导') }}</td>
          </tr>
        </tbody>
      </table>
      <table m="t-4" text="xl sm">
        <thead>
          <tr bg="blue-700">
            <th p="2" font="normal" border="~">劝化</th>
            <th p="2" font="normal" border="~">辨别诸灵</th>
            <th p="2" font="normal" border="~">施舍</th>
            <th p="2" font="normal" border="~">帮助人</th>
            <th p="2" font="normal" border="~">怜悯</th>
          </tr>
        </thead>
        <tbody>
          <tr bg="blue-600">
            <td p="2" text="center" border="~">{{ getGift('劝化') }}</td>
            <td p="2" text="center" border="~">{{ getGift('辨别诸灵') }}</td>
            <td p="2" text="center" border="~">{{ getGift('施舍') }}</td>
            <td p="2" text="center" border="~">{{ getGift('帮助人') }}</td>
            <td p="2" text="center" border="~">{{ getGift('怜悯') }}</td>
          </tr>
        </tbody>
      </table>
      <table m="t-4" text="xl sm">
        <thead>
          <tr bg="blue-700">
            <th p="2" font="normal" border="~">使徒</th>
            <th p="2" font="normal" border="~">单身</th>
            <th p="2" font="normal" border="~">代祷</th>
            <th p="2" font="normal" border="~">舍命受苦(殉道)</th>
            <th p="2" font="normal" border="~">服事</th>
          </tr>
        </thead>
        <tbody>
          <tr bg="blue-600">
            <td p="2" text="center" border="~">{{ getGift('使徒') }}</td>
            <td p="2" text="center" border="~">{{ getGift('单身') }}</td>
            <td p="2" text="center" border="~">{{ getGift('代祷') }}</td>
            <td p="2" text="center" border="~">{{ getGift('舍命受苦(殉道)') }}</td>
            <td p="2" text="center" border="~">{{ getGift('服事') }}</td>
          </tr>
        </tbody>
      </table>
      <table m="t-4" text="xl sm">
        <thead>
          <tr bg="blue-700">
            <th p="2" font="normal" border="~">治理</th>
            <th p="2" font="normal" border="~">神迹奇事情</th>
            <th p="2" font="normal" border="~">医治</th>
            <th p="2" font="normal" border="~">说方言</th>
            <th p="2" font="normal" border="~">翻方言</th>
          </tr>
        </thead>
        <tbody>
          <tr bg="blue-600">
            <td p="2" text="center" border="~">{{ getGift('治理') }}</td>
            <td p="2" text="center" border="~">{{ getGift('神迹奇事') }}</td>
            <td p="2" text="center" border="~">{{ getGift('医治') }}</td>
            <td p="2" text="center" border="~">{{ getGift('说方言') }}</td>
            <td p="2" text="center" border="~">{{ getGift('翻方言') }}</td>
          </tr>
        </tbody>
      </table>
      <table m="t-4" text="xl sm">
        <thead>
          <tr bg="blue-700">
            <th p="2" font="normal" border="~">先知</th>
            <th p="2" font="normal" border="~">牧师</th>
            <th p="2" font="normal" border="~">教导</th>
            <th p="2" font="normal" border="~">智慧的言语</th>
            <th p="2" font="normal" border="~">知识的言语</th>
          </tr>
        </thead>
        <tbody>
          <tr bg="blue-600">
            <td p="2" text="center" border="~">{{ getGift('先知') }}</td>
            <td p="2" text="center" border="~">{{ getGift('牧师') }}</td>
            <td p="2" text="center" border="~">{{ getGift('教导') }}</td>
            <td p="2" text="center" border="~">{{ getGift('智慧的言语') }}</td>
            <td p="2" text="center" border="~">{{ getGift('知识的言语') }}</td>
          </tr>
        </tbody>
      </table>
      <p m="t-2" text="sm center">
        {{ t('test.gift.description') }}
        <span font="bold">{{ getHighestGiftLabel() }}</span>
      </p>
    </template>
    <div v-else-if="giftResponses && user && !gift" flex="~ col" align="items-center">
      <p flex="~" align="items-center">
        {{ t('test.gift.no_test_done') }}
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
        @click="push('/test/gift')"
      >
        {{ t('test.gift.button') }}
      </button>
    </div>
    <Spinner v-else animate="spin" m="y-8 x-auto" w="12" h="12" text="neonGreen" />
  </div>
</template>
