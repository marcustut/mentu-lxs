<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { RadioGroup, RadioGroupLabel, RadioGroupOption } from '@headlessui/vue';
import { disc as data } from '~/data';
import { firebase, db } from '~/modules/firebase';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import type { Disc } from '~/types';

const discResponsesRef = db.collection('disc-responses');

// @ts-ignore
const discResponses = useFirestore<Disc[]>(discResponsesRef);
const { user } = useAuth(firebase.auth());
const { push } = useRouter();
const { t } = useI18n();

const disc = ref<Disc>();
// @ts-ignore
const answers = ref<Disc['responses']>({});

watch(discResponses, () => {
  if (!discResponses.value) return;
  if (!user.value) return;

  disc.value = discResponses.value.find((r) => r.uid === user.value?.uid);
});

watch(disc, () => {
  if (!disc.value) return;

  if (disc.value) {
    alert('You have already done the test\nYou will be redirected to the result page');
    push('/test');
  }
});

const submit = () => {
  if (!user.value) return;
  if (Object.keys(answers.value).length !== 40) {
    alert('Please finish all 40 questions');
    return;
  }

  discResponsesRef
    .add({
      uid: user.value.uid,
      displayName: user.value.displayName,
      responses: answers.value,
    } as Disc)
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
    <h1 text="2xl" font="bold">{{ t('test.disc.title') }}</h1>

    <RadioGroup
      v-for="({ question }, index) in data"
      v-model="answers[index]"
      :key="index"
      :m="index !== 0 ? 't-8' : ''"
      w="full"
    >
      <RadioGroupLabel text="lg" font="bold"
        >{{
          // @ts-ignore
          parseInt(index) + 1
        }}. 你是一个怎样的人？</RadioGroupLabel
      >
      <RadioGroupOption
        v-for="ans in Object.entries(question)"
        v-slot="{ checked }"
        :key="ans[1]"
        :value="ans[0]"
      >
        <button
          m="t-2"
          p="x-4 y-2"
          :bg="`${checked ? 'blue-500' : 'gray-200 dark:dark-400'}`"
          text="left sm"
          border="rounded-2xl"
          outline="focus:none"
          transition="colors duration-200 ease-in-out"
        >
          {{ ans[1] }}
        </button>
      </RadioGroupOption>
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
      {{ t('test.disc.submit') }}
    </button>
  </div>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
