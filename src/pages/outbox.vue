<script setup lang="ts">
import { useAuth, useFirestore } from '@vueuse/firebase';
import { Form, Field, ErrorMessage } from 'vee-validate';
import { useI18n } from 'vue-i18n';
import { NUpload, NUploadDragger, NP, NText, NIcon } from 'naive-ui';
import { firebase, db, storage } from '~/modules/firebase';
import { isDark } from '~/logic/dark';
import Spinner from '~/components/Spinner.vue';
import * as Yup from 'yup';
import type { User } from '~/types';
import type { Ref } from 'vue';

type FileInfo = {
  id: string;
  name: string;
  url: string | null;
  percentage: number;
  status: 'pending' | 'uploading' | 'finished' | 'removed' | 'error';
  file: File | null;
};

const schema = Yup.object({
  receiver: Yup.string().required('receiver must be specified'),
  message: Yup.string().required('message cannot be empty'),
});

const usersRef = db.collection('users');
const messagesRef = db.collection('messages');

const { isAuthenticated, user: firebaseUser } = useAuth(firebase.auth());

const files = ref<File[]>([]);
const user = ref<User | undefined>();
// @ts-ignore
const users = useFirestore<User[]>(usersRef);

watch(firebaseUser, () => {
  if (!firebaseUser.value) return;

  const userRef = usersRef.where('uid', '==', firebaseUser.value.uid);
  userRef.get().then((query) => {
    if (query.size === 0 || query.size > 1) {
      alert('Something wrong happened, please contact support');
      return;
    }
    query.forEach((doc) => {
      user.value = doc.data() as User;
    });
  });
});

const handleUpload = ({
  file,
  fileList,
}: {
  file: FileInfo;
  fileList: FileInfo[];
  event: ProgressEvent | Event | undefined;
}) => {
  if (!file.file) {
    console.error('Upload file error');
    return;
  }

  const found = files.value.find((f) => f.name === file.file?.name);

  if (found) {
    fileList.pop();
    return;
  }

  files.value.push(file.file);
};

type FormSubmitParam = {
  receiver: string;
  message: string;
};

const handleSubmit = async ({ receiver, message }: FormSubmitParam) => {
  if (!firebaseUser.value) return;
  if (!receiver) return;
  if (firebaseUser.value.uid == receiver) {
    alert('Cannot send message to yourself');
    return;
  }

  const attachment: string[] = [];

  const messagePath = `Messages/${receiver}`;

  for (let i = 0; i < files.value.length; i++) {
    const file = files.value[i];
    const fileRef = storage.ref(`${messagePath}/${file.name}`);
    await fileRef.put(file);
    const url = await fileRef.getDownloadURL();
    attachment.push(url);
  }

  const messageObj = {
    attachment,
    message,
    receiver,
    sender: firebaseUser.value.uid,
    created_at: new Date(),
  };

  await messagesRef.add(messageObj);

  alert(`Your message is successfully sent!`);
};

const capitalize = (s: string | undefined) => {
  if (!s) return '';

  return s
    .split(' ')
    .map((a) => {
      if (!a) return '';
      return a[0].toUpperCase() + a.slice(1);
    })
    .join(' ');
};

const { locale } = useI18n() as unknown as { locale: Ref<'en' | 'zh-CN'> };
</script>

<template>
  <template v-if="isAuthenticated">
    <div m="x-auto" w="full sm:[500px]">
      <div flex="~" align="items-center">
        <h3 font="bold" text="3xl">Your Outbox</h3>
        <twemoji:outbox-tray w="7" h="7" m="l-3" />
      </div>

      <Form
        v-slot="{ isSubmitting }"
        :validate-on-mount="false"
        :validation-schema="schema"
        @submit="handleSubmit"
        m="t-4"
        w="full"
        flex="~ col"
        align="items-center"
      >
        <div
          v-if="user"
          w="full"
          p="x-3 y-2"
          bg="gray-200 dark:dark-300"
          text="sm"
          border="rounded-2xl"
          flex="~"
          align="items-center"
        >
          <p bg="neonGreen" border="rounded-2xl" p="x-2 y-1" text="textDark xs" font="medium">
            Sender
          </p>
          <input
            name="sender"
            :placeholder="`${
              user.roles.includes('trainee') ? user.trainee_id : user.roles[0].toUpperCase()
            } - ${locale === 'en' ? user.name.en : user.name['zh-CN']}`"
            m="l-2"
            w="full"
            bg="gray-200 dark:dark-300"
            text="sm placeholder-gray-700 dark:placeholder-gray-200"
            font="bold"
            outline="focus:none"
            disabled
          />
        </div>
        <div
          w="full"
          m="t-4"
          p="x-3 y-2"
          bg="gray-200 dark:dark-300"
          text="sm"
          border="rounded-2xl"
          flex="~"
          align="items-center"
        >
          <p bg="neonGreen" border="rounded-2xl" p="x-2 y-1" text="textDark xs" font="medium">
            Receiver
          </p>
          <Field
            v-if="users"
            as="select"
            name="receiver"
            m="l-2"
            w="full"
            bg="gray-200 dark:dark-300"
            text="sm gray-700 dark:gray-200 placeholder-gray-700 dark:placeholder-gray-200 capitalize"
            font="bold"
            outline="focus:none"
          >
            <option v-for="user in users" :key="user.uid" :value="user.uid">
              {{ capitalize(user.name.en?.toLowerCase()) }}
            </option>
          </Field>
          <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
        </div>
        <ErrorMessage name="receiver" m="t-1" text="red-500 sm" />

        <Field
          as="textarea"
          name="message"
          placeholder="Enter you message here..."
          w="full"
          h="[300px]"
          m="t-4"
          p="4"
          bg="gray-200 dark:dark-300"
          text="sm textDark dark:white"
          border="rounded-2xl"
          outline="focus:none"
          style="resize: none"
        />
        <ErrorMessage name="message" m="t-1" text="red-500 sm" />

        <n-upload
          name="upload"
          m="t-4"
          w="full"
          accept="image/*, audio/*, video/*"
          @change="handleUpload"
        >
          <n-upload-dragger :style="isDark && 'background: #2d2d2d'">
            <div style="margin-bottom: 12px">
              <n-icon size="48" :depth="3">
                <feather:upload w="8" h="8" text="dark:white" />
              </n-icon>
            </div>
            <n-text text="dark:white">Click or drag file to this area to upload</n-text>
            <n-p depth="3" style="margin: 6px 0 0 0; font-size: 11px"
              >Only upload either video / photo / audio</n-p
            >
          </n-upload-dragger>
        </n-upload>
        <button
          w="full"
          p="x-4 y-3"
          m="t-4"
          bg="gray-200 dark:dark-300"
          text="sm textDark dark:white"
          ring="hover:2"
          font="medium"
          flex="~"
          justify="center"
          align="items-center"
          border="rounded-2xl"
          outline="focus:none"
        >
          <template v-if="!isSubmitting">
            Submit <twemoji:open-mailbox-with-raised-flag m="l-2" w="5" h="5" />
          </template>
          <Spinner v-else animate="spin" m="x-auto" w="6" h="6" text="neonGreen" />
        </button>
      </Form>
    </div>
  </template>
  <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
</template>

<route lang="yaml">
meta:
  layout: app
</route>
