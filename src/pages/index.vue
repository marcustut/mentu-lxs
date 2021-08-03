<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { useIntervalFn } from '@vueuse/shared';
import { Icon } from '@iconify/vue';
import { useI18n } from 'vue-i18n';
import type { Ref } from 'vue';
import { db, firebase } from '~/modules/firebase';
import Spinner from '~/components/Spinner.vue';
import type { Notification, Event, User } from '~/types';

const greetings = ['Hello', 'Hi', 'Yo', 'Hey', 'Hola', 'こんにちは', 'Bonjour', 'Salut', '你好'];
const weekdays: Record<number, string> = {
  0: 'Sunday',
  1: 'Monday',
  2: 'Tuesday',
  3: 'Wednesday',
  4: 'Thursday',
  5: 'Friday',
  6: 'Saturday',
};
const weekdays_CN: Record<number, string> = {
  0: '星期日',
  1: '星期一',
  2: '星期二',
  3: '星期三',
  4: '星期四',
  5: '星期五',
  6: '星期六',
};
const roles = [
  { id: 1, name: { en: 'Trainee', 'zh-CN': '学员' }, value: 'trainee' },
  { id: 2, name: { en: 'Producer', 'zh-CN': '制作人' }, value: 'producer' },
  { id: 3, name: { en: 'Mentor', 'zh-CN': '导师' }, value: 'mentor' },
  { id: 4, name: { en: 'WM', 'zh-CN': '工作人员' }, value: 'wm' },
  { id: 5, name: { en: 'Production Crew', 'zh-CN': '制作团队' }, value: 'crew' },
];

const notificationsRef = db.collection('notifications');
const eventsRef = db.collection('events');
const usersRef = db.collection('users');

// @ts-ignore
const notifications = useFirestore<Notification[]>(notificationsRef);
// @ts-ignore
const events = useFirestore<Event[]>(eventsRef);
// @ts-ignore
const users = useFirestore<User[]>(usersRef);

const { user: firebaseUser } = useAuth(firebase.auth());

const user = ref<User | undefined>();

watch(users, () => {
  if (!users.value) return;
  if (!firebaseUser.value) return;

  user.value = users.value.find((u) => u.uid === firebaseUser.value?.uid);
});

const { locale } = useI18n() as unknown as { locale: Ref<string> };
const { t } = useI18n();

const word = ref('Hello');
const event = ref<Event | null>(null);

watch(events, () => {
  if (!events.value) return;

  event.value = events.value.sort(
    (a, b) => b.start_datetime.toDate().getTime() - a.start_datetime.toDate().getTime()
  )[0];
});

const { pause, resume, isActive } = useIntervalFn(() => {
  word.value = greetings[Math.round(Math.random() * (greetings.length - 1))];
}, 500);

const greetingHandler = () => {
  if (isActive.value) pause();
  else resume();
};
</script>

<template>
  <div
    v-if="user"
    w="full"
    h="2/7"
    gradient="to-r from-neonGreen to-emerald-400"
    pos="absolute top-0 left-0"
    border="rounded-b-3xl"
    transition="~ duration-200 ease-in-out"
    :style="{ zIndex: -1 }"
  />

  <div v-if="user">
    <div text="space-pre-wrap 3xl gray-700" font="bold" @click="greetingHandler">
      <h2 text="xl">{{ word }},</h2>
      <h1>
        {{ locale === 'en' ? user.name.en : user.name['zh-CN'] }}
      </h1>
    </div>

    <div
      v-if="user"
      m="t-8"
      p="x-4 y-3"
      bg="gray-100"
      flex="~ col"
      text="gray-700"
      border="rounded-2xl"
      shadow="xl"
    >
      <div flex="~" align="items-center">
        <!-- <div
          m="l-2"
          p="x-2"
          h="5"
          bg="neonGreen"
          border="rounded-3xl"
          text="xs"
          flex="~"
          justify="center"
          align="items-center"
        >
          {{ t('home.identification.trainee') }}
        </div> -->
        <div flex="~ wrap" align="items-center">
          <h3 m="r-2" text="xl" font="semibold">
            {{ locale === 'en' ? user.name.en : user.name['zh-CN'] }}
          </h3>
          <div
            v-for="role in user.roles"
            :key="role"
            p="x-2"
            h="5"
            m="r-2"
            border="rounded-3xl"
            text="xs gray-700"
            flex="~"
            justify="center"
            align="items-center"
            :class="{
              'bg-neonGreen': role === 'trainee',
              'bg-yellow-300': role === 'producer',
              'bg-blue-300': role === 'mentor',
              'bg-green-300': role === 'wm',
              'bg-violet-200': role === 'crew',
            }"
          >
            {{
              locale === 'en'
                ? roles.find((r) => r.value === role)?.name.en
                : roles.find((r) => r.value === role)?.name['zh-CN']
            }}
          </div>
        </div>
      </div>
      <div m="t-1" flex="~ col" text="sm" font="medium">
        <div v-show="user.trainee_id" flex="~" justify="between">
          <p>{{ t('home.identification.id_number_label') }} :</p>
          <p>{{ user.trainee_id }}</p>
        </div>
        <div flex="~" justify="between">
          <p>{{ t('home.identification.programme_label') }} :</p>
          <p>{{ t('home.identification.programme') }}</p>
        </div>
        <div flex="~" justify="between">
          <p>{{ t('home.identification.organized_by_label') }} :</p>
          <p>{{ t('home.identification.organized_by') }}</p>
        </div>
      </div>
    </div>

    <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />

    <div v-if="event" m="t-8">
      <h3 text="xl" font="medium">
        {{ t('home.schedule') }}
      </h3>
      <div
        m="t-4"
        p="4"
        gradient="to-r from-neonGreen to-emerald-300"
        shadow="lg"
        border="rounded-2xl"
        flex="~ col"
        text="gray-700"
      >
        <a flex="~" align="items-center" target="_blank" :href="event.title_link">
          <div bg="neonGreenShade" p="2" border="rounded-xl">
            <Icon :icon="event.icon" w="10" h="10" />
          </div>
          <div m="l-4">
            <p text="lg" font="bold">
              {{ locale === 'en' ? event.title.en : event.title['zh-CN'] }}
            </p>
            <p text="sm">
              {{
                `${event.start_datetime.toDate().toLocaleDateString()} (${
                  locale === 'en'
                    ? weekdays[event.start_datetime.toDate().getDay()]
                    : weekdays_CN[event.start_datetime.toDate().getDay()]
                })`
              }}
            </p>
            <p text="sm" font="medium">
              {{ event.start_datetime.toDate().toLocaleTimeString() }}
            </p>
          </div>
        </a>
        <div m="y-3" w="full" h="[1px]" bg="gray-700" />
        <div flex="~" align="items-center" text="sm">
          <div bg="neonGreenShade" m="r-2" p="2" border="rounded-xl">
            <twemoji-light-bulb w="6" h="6" />
          </div>
          <a target="_blank" :href="event.description_link">
            {{ locale === 'en' ? event.description.en : event.description['zh-CN'] }}
          </a>
        </div>
      </div>
    </div>

    <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />

    <div v-if="notifications" m="t-8">
      <h3 m="b-4" text="xl" font="medium">
        {{ t('home.notification') }}
      </h3>
      <a
        v-for="(notification, index) in notifications.filter((n) => n.active)"
        :key="notification.title.en"
        :m="`${index === 0 ? 't-4' : ''} ${index !== notifications.length - 1 ? 'b-4' : ''}`"
        :href="notification.link"
        target="_blank"
        p="4"
        bg="gray-100"
        flex="~"
        align="items-center"
        border="rounded-2xl"
        class="text-gray-700"
      >
        <div bg="neonGreen" p="2" border="rounded-lg">
          <Icon :icon="notification.icon" w="6" h="6" />
        </div>
        <div m="l-4" flex="~ col">
          <p font="bold">
            {{ locale === 'en' ? notification.title.en : notification.title['zh-CN'] }}
          </p>
          <p text="xs">
            {{ locale === 'en' ? notification.description.en : notification.description['zh-CN'] }}
          </p>
        </div>
        <bi-arrow-up-right-circle m="l-auto" />
      </a>
    </div>

    <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
  </div>

  <div
    v-else
    m="x-auto"
    h="screen"
    w="screen"
    pos="fixed top-0 left-0"
    flex="~"
    justify="center"
    align="items-center"
  >
    <Spinner animate="spin" w="12" h="12" text="neonGreen" />
  </div>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
