<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { RadioGroup, RadioGroupLabel, RadioGroupOption } from '@headlessui/vue';
import { gift as data, giftEn as dataEn } from '~/data';
import { firebase, db } from '~/modules/firebase';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import type { Gift } from '~/types';

const giftResponsesRef = db.collection('gift-responses');

// @ts-ignore
const giftResponses = useFirestore<Gift[]>(giftResponsesRef);
const { user } = useAuth(firebase.auth());
const { push } = useRouter();
const { t, locale } = useI18n();

const gift = ref<Gift>();
// @ts-ignore
const answers = ref<Gift['responses']>({});

const questions = ref<typeof data>(dataEn);

watch(locale, () => {
  if (locale.value === 'en') {
    questions.value = dataEn;
  } else {
    questions.value = data;
  }
});

watch(giftResponses, () => {
  if (!giftResponses.value) return;
  if (!user.value) return;

  gift.value = giftResponses.value.find((r) => r.uid === user.value?.uid);
});

watch(gift, () => {
  if (!gift.value) return;

  if (gift.value) {
    alert('You have already done the test\nYou will be redirected to the result page');
    push('/test');
  }
});

const submit = () => {
  if (!user.value) return;
  if (Object.keys(answers.value).length !== Object.keys(data).length) {
    console.log(Object.keys(answers.value).length);
    console.log(Object.keys(data).length);
    alert('Please finish all 125 questions');
    return;
  }

  giftResponsesRef
    .add({
      uid: user.value.uid,
      displayName: user.value.displayName,
      responses: answers.value,
    } as Gift)
    .then(() => {
      alert('Your test is submitted succesfully 🎉\nYou will be redirected to the results page');
      push('/test');
    })
    .catch((err) =>
      alert('An error had occured, please contact our production crew\nError Message: ' + err)
    );
};
</script>

<template>
  <div p="b-4" flex="~ col" align="items-center">
    <h1 text="2xl" font="bold">{{ t('test.gift.title') }}</h1>

    <RadioGroup
      v-for="(question, index) in questions"
      v-model="answers[index]"
      :key="index"
      :m="index !== 0 ? 't-8' : ''"
      w="full"
      p="4"
      bg="gray-200 dark:dark-400"
      border="rounded-2xl"
    >
      <RadioGroupLabel text="lg" font="bold">{{
        // @ts-ignore
        `${parseInt(index) + 1}. ${question}`
      }}</RadioGroupLabel>
      <div m="t-4 b-2" w="full" flex="~" justify="around" align="items-center">
        <p m="r-4" text="xs">{{ t('test.gift.totally_no') }}</p>
        <RadioGroupOption
          v-for="ans in Array.from([0, 1, 2, 3, 4])"
          v-slot="{ checked }"
          :key="ans"
          :value="ans"
        >
          <button
            m="x-1"
            p="x-3 y-1"
            :bg="`${checked ? 'blue-500' : 'gray-300 dark:dark-500'}`"
            :text="`left sm ${checked ? 'gray-50' : ''}`"
            border="rounded-2xl"
            outline="focus:none"
            transition="colors duration-200 ease-in-out"
            name="clickme"
          >
            {{ ans }}
          </button>
        </RadioGroupOption>
        <p m="l-4" text="xs">{{ t('test.gift.always_yes') }}</p>
      </div>
    </RadioGroup>

    <button
      m="t-4 b-1"
      p="x-4 y-3"
      bg="neonGreen"
      text="gray-700 sm"
      font="medium"
      border="rounded-2xl"
      outline="focus:none"
      @click="submit"
    >
      {{ t('test.gift.submit') }}
    </button>
  </div>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
