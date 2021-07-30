<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase'
import { useIntervalFn } from '@vueuse/shared'
import { Icon } from '@iconify/vue'
import { useI18n } from 'vue-i18n'
import type { Ref } from 'vue'
import { db, firebase } from '~/modules/firebase'
import Spinner from '~/components/Spinner.vue'
import type { Notification, Event } from '~/types'

const greetings = ['Hello', 'Hi', 'Yo', 'Hey', 'Hola', 'こんにちは', 'Bonjour', 'Salut', '你好']
const weekdays = {
  0: 'Sunday',
  1: 'Monday',
  2: 'Tuesday',
  3: 'Wednesday',
  4: 'Thursday',
  5: 'Friday',
  6: 'Saturday',
} as Record<number, string>
const weekdays_CN = {
  0: '星期日',
  1: '星期一',
  2: '星期二',
  3: '星期三',
  4: '星期四',
  5: '星期五',
  6: '星期六',
} as Record<number, string>

const notificationsRef = db.collection('notifications')
const eventsRef = db.collection('events')

// @ts-ignore
const notifications = useFirestore<Notification[]>(notificationsRef)
// @ts-ignore
const events = useFirestore<Event[]>(eventsRef)
const { user } = useAuth(firebase.auth())
const { locale } = useI18n() as unknown as { locale: Ref<string> }
const { t } = useI18n()

const word = ref('Hello')
const event = ref<Event | null>(null)

watch(events, () => {
  if (!events.value) return

  event.value = events.value.sort((a, b) =>
    a.updatedAt.toDate().getTime() - b.updatedAt.toDate().getTime())[0]
})

const { pause, resume, isActive } = useIntervalFn(() => {
  word.value = greetings[Math.round(Math.random() * (greetings.length - 1))]
}, 500)

const greetingHandler = () => {
  if (isActive.value) pause()
  else resume()
}
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
      <h2 text="xl">
        {{ word }},
      </h2>
      <h1>
        {{ `${user?.displayName}` }}
      </h1>
    </div>

    <div
      m="t-8"
      p="x-4 y-3"
      bg="gray-100"
      flex="~ col"
      text="gray-700"
      border="rounded-2xl"
      shadow="xl"
    >
      <div flex="~" align="items-center">
        <h3 text="xl" font="semibold">
          {{ user?.displayName }}
        </h3>
        <div
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
        </div>
      </div>
      <div m="t-1" flex="~ col" text="sm" font="medium">
        <div flex="~" justify="between">
          <p>{{ t('home.identification.id_number_label') }} :</p>
          <p>D217001</p>
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
              {{ `${event.updatedAt.toDate().toLocaleDateString()} (${locale === 'en' ? weekdays[event.updatedAt.toDate().getDay()] : weekdays_CN[event.updatedAt.toDate().getDay()]})` }}
            </p>
            <p text="sm" font="medium">
              {{ event.updatedAt.toDate().toLocaleTimeString() }}
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

    <Spinner
      v-else
      animate="spin"
      m="t-8 x-auto"
      w="12"
      h="12"
      text="neonGreen"
    />

    <div v-if="notifications" m="t-8">
      <h3 text="xl" font="medium">
        {{ t('home.notification') }}
      </h3>
      <a
        v-for="(notification, index) in notifications"
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

    <Spinner
      v-else
      animate="spin"
      m="t-8 x-auto"
      w="12"
      h="12"
      text="neonGreen"
    />
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
