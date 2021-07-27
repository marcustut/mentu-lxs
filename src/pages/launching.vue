<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase'
import vueDanmaku from 'vue3-danmaku'
import { firebase, db } from '~/modules/firebase'
import Spinner from '~/components/Spinner.vue'

const { user } = useAuth(firebase.auth())

const messageText = ref('')
const danmus = ref([{ avatar: 'http://a.com/a.jpg', name: 'a', text: 'aaa' }, { avatar: 'http://a.com/b.jpg', name: 'b', text: 'bbb' }])

const launchingRef = db.collection('interaction').doc('launching')
const messagesRef = db.collection('interaction').doc('launching').collection('messages')

const launching = useFirestore(launchingRef)
const messages = useFirestore(messagesRef)

const addCount = () => {
  launchingRef.update({
    count: firebase.firestore.FieldValue.increment(1),
  })
}

const sendMessage = () => {
  if (messageText.value !== '') {
    messagesRef.add({
      content: messageText.value,
      sentBy: { name: user.value?.displayName, avatar_url: user.value?.photoURL },
    })
    messageText.value = ''
  }
}

// db.collection('interaction').where('name', '==', 'launching').get().then(querySnapshot => querySnapshot.forEach(doc => console.log(doc.data())))
// db.collection('interaction').doc('launching').get()
</script>

<template>
  <button
    p="x-4 y-3"
    z="10"
    pos="fixed bottom-4 right-4"
    bg="neonGreen"
    border="rounded-2xl"
    shadow="xl"
    text="xl"
    outline="focus:none"
    @click="addCount"
  >🎉</button>
  <div h="[91vh]" flex="~ col" justify="center" align="items-center" pos="relative">
    <template v-if="launching">
      <div h="2" w="72" pos="relative">
        <div
          h="full"
          bg="neonGreen"
          border="rounded"
          pos="absolute"
          :style="{ width: `${launching.count}%` }"
        ></div>
        <div h="full" w="full" bg="neonGreenShade" border="rounded"></div>
      </div>

      <input
        v-model="messageText"
        placeholder="Enter your messsage here! 🥺"
        m="y-4"
        w="2/3 sm:1/3"
        p="x-4 y-3"
        bg="true-gray-50 dark:true-gray-700"
        text="sm"
        border="rounded-2xl"
        shadow="lg"
        outline="focus:none"
      />
      <button
        p="x-4 y-3"
        bg="neonGreen"
        text="sm dark-700"
        font="medium"
        flex="~"
        align="items-center"
        border="rounded-2xl"
        outline="focus:none"
        shadow="lg"
        @click="sendMessage"
      >
        Send
        <bx-bx-mail-send w="4" h="4" m="l-1" />
      </button>

      <div v-if="messages" pos="absolute top-0">
        <vue-danmaku :danmus="messages" use-slot w="screen" h="[250px]">
          <template #dm="{ danmu }">
            <div flex="~" align="items-center">
              <img
                border="rounded-full"
                w="6"
                h="6"
                m="r-1"
                :src="danmu.sentBy.avatar_url"
                :alt="danmu.sentBy.name"
              />
              <span text="gray-700 dark:gray-200 sm">{{ danmu.content }}</span>
            </div>
          </template>
        </vue-danmaku>
      </div>
    </template>
    <Spinner v-else w="12" h="12" text="neonGreen" />
  </div>
</template>

<route lang="yaml">
meta:
  layout: app
</route>
