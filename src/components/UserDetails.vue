<script setup lang="ts">
import { Dialog, DialogOverlay, Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue';
import { useAuth, useFirestore } from '@vueuse/firebase';
import { useForm, useField } from 'vee-validate';
import { firebase, db } from '~/modules/firebase';
import * as Yup from 'yup';
import { useTimeout } from '@vueuse/core';
import { useI18n } from 'vue-i18n';
import type { User } from '~/types';
import type { Ref } from 'vue';

const roles = [
  { id: 1, name: { en: 'Trainee', 'zh-CN': '学员' }, value: 'trainee' },
  { id: 2, name: { en: 'Producer', 'zh-CN': '制作人' }, value: 'producer' },
  { id: 3, name: { en: 'Mentor', 'zh-CN': '导师' }, value: 'mentor' },
  { id: 4, name: { en: 'WM', 'zh-CN': '工作人员' }, value: 'wm' },
  { id: 5, name: { en: 'Production Crew', 'zh-CN': '制作团队' }, value: 'crew' },
];

const selectedRoles = ref([roles[0]]);

const { locale } = useI18n() as unknown as { locale: Ref<string> };
const { t } = useI18n();

const usersRef = db.collection('users');
const schema = Yup.object({
  name: Yup.string()
    .required(t('user_details.form.error.name'))
    .min(3, 'Name must be at least 3 characters'),
  name_cn: Yup.string()
    .required(t('user_details.form.error.name_cn'))
    .min(2, 'Chinese name must be at least 2 chracters'),
  age: Yup.number()
    .required(t('user_details.form.error.age'))
    .min(12, 'You must be at least 12 years old'),
  phone_number: Yup.string()
    .required(t('user_details.form.error.phone_number'))
    .matches(/^(\+?6?01)[0|1|2|3|4|6|7|8|9]-*[0-9]{7,8}$/, 'Not a valid Malaysian phone number'),
  trainee_id: Yup.string().optional(),
});

const { handleSubmit, isSubmitting } = useForm({ validationSchema: schema });

const { value: name, errorMessage: nameError } = useField('name');
const { value: name_cn, errorMessage: nameCnError } = useField('name_cn');
const { value: age, errorMessage: ageError } = useField('age');
const { value: phone_number, errorMessage: phoneNumberError } = useField('phone_number');
const { value: trainee_id, errorMessage: traineeIdError } = useField('trainee_id');

// @ts-ignore
const users = useFirestore<User[]>(usersRef);

const { isAuthenticated, user: firebaseUser } = useAuth(firebase.auth());

const user = ref<User | undefined>();

// Timeout to wait for 'isAuthenticated'
const ready = useTimeout(500);

watch(users, () => {
  if (!users.value) return;
  if (!firebaseUser.value) return;

  user.value = users.value.find((u) => u.uid === firebaseUser.value?.uid);
});

const submitUserDetails = handleSubmit((values) => {
  // If not logged in
  if (!firebaseUser.value) return;

  // If no role selected
  if (selectedRoles.value.length === 0) {
    alert(t('user_details.form.error.roles'));
    return;
  }

  // If is trainee but no id
  if (selectedRoles.value.find((r) => r.value === 'trainee') && !values.trainee_id) {
    alert(t('user_details.form.error.trainee_id'));
    return;
  }

  const user: User = {
    uid: firebaseUser.value.uid,
    trainee_id: values.trainee_id ? values.trainee_id : null,
    displayName: firebaseUser.value.displayName ? firebaseUser.value.displayName : null,
    email: firebaseUser.value.email ? firebaseUser.value.email : null,
    name: { en: values.name ? values.name : null, 'zh-CN': values.name_cn ? values.name_cn : null },
    roles: selectedRoles.value.map((r) => r.value),
    age: values.age ? values.age : null,
    phone_number: values.phone_number ? values.phone_number : null,
  };

  usersRef.add(user);
});

const removeRole = (role: string) => {
  if (!roles || roles.length === 0) return;
  selectedRoles.value = selectedRoles.value.filter((r) => r.value !== role);
};

const addRole = (role: string) => {
  // If list of roles is not defined
  if (!roles || roles.length === 0) return;

  // If this role already exists
  if (selectedRoles.value.find((r) => r.value === role)) return;

  // Get the role to add
  const roleToAdd = roles.find((r) => r.value === role);

  // If unable to get the role
  if (!roleToAdd) return;

  // Push the role to selectedRoles
  selectedRoles.value.push(roleToAdd);
};
</script>

<template>
  <Dialog :open="ready && isAuthenticated && !user">
    <DialogOverlay pos="fixed inset-0" backdrop="~ blur-[2px]" />

    <div
      w="<sm:4/5"
      p="x-6 y-4"
      bg="gray-50 dark:dark-400"
      pos="fixed top-[50%] left-[50%]"
      flex="~ col"
      align="items-center"
      text="gray-700 dark:gray-200"
      border="rounded-3xl"
      shadow="2xl"
      transform="~ translate-x-[-50%] translate-y-[-50%]"
      transition="duration-200 ease-in-out"
    >
      <template v-if="!isSubmitting">
        <div flex="~" align="items-center">
          <h4 text="2xl" font="bold">{{ t('user_details.header') }}</h4>
          <twemoji-clipboard w="6" h="6" m="l-2" />
        </div>
        <p m="t-2" text="sm center" font="leading-tight">{{ t('user_details.description') }}</p>

        <form m="t-4" w="full" flex="~ col" align="items-center" @submit="submitUserDetails">
          <input
            name="name"
            v-model="name"
            :placeholder="t('user_details.form.name')"
            w="full"
            p="x-3 y-2"
            bg="gray-200 dark:dark-300"
            text="sm"
            border="rounded-2xl"
            outline="focus:none"
          />
          <span m="t-2" text="xs red-500">{{ nameError }}</span>

          <input
            name="name_cn"
            v-model="name_cn"
            :placeholder="t('user_details.form.name_cn')"
            m="t-2"
            w="full"
            p="x-3 y-2"
            bg="gray-200 dark:dark-300"
            text="sm"
            border="rounded-2xl"
            outline="focus:none"
          />
          <span m="t-2" text="xs red-500">{{ nameCnError }}</span>

          <input
            name="phone_number"
            v-model="phone_number"
            :placeholder="t('user_details.form.phone_number')"
            m="t-2"
            w="full"
            p="x-3 y-2"
            bg="gray-200 dark:dark-300"
            text="sm"
            border="rounded-2xl"
            outline="focus:none"
          />
          <span m="t-2" text="xs red-500">{{ phoneNumberError }}</span>

          <input
            name="age"
            v-model="age"
            :placeholder="t('user_details.form.age')"
            type="number"
            m="t-2"
            w="full"
            p="x-3 y-2"
            bg="gray-200 dark:dark-300"
            text="sm"
            border="rounded-2xl"
            outline="focus:none"
          />
          <span m="t-2" text="xs red-500">{{ ageError }}</span>

          <div m="y-4" w="full" h="[1px]" bg="gray-300 dark:dark-200" border="rounded-full" />

          <div align="self-start" flex="~ wrap">
            <button
              v-for="role in selectedRoles"
              :key="role.id"
              p="x-2"
              h="5"
              m="r-2 t-2"
              border="rounded-3xl"
              text="xs gray-700"
              flex="~"
              justify="center"
              align="items-center"
              outline="focus:none"
              :class="{
                'bg-neonGreen': role.value === 'trainee',
                'bg-yellow-300': role.value === 'producer',
                'bg-blue-300': role.value === 'mentor',
                'bg-green-300': role.value === 'wm',
                'bg-violet-200': role.value === 'crew',
              }"
              @click="
                (e) => {
                  e.preventDefault();
                  removeRole(role.value);
                }
              "
            >
              {{ locale === 'en' ? role.name.en : role.name['zh-CN'] }}
              <iconoir-cancel m="l-[2px]" w="3" h="3" />
            </button>
          </div>

          <Menu>
            <MenuButton
              :m="selectedRoles.length !== 0 ? 't-6' : 't-2'"
              w="full"
              p="x-3 y-2"
              bg="gray-200 dark:dark-300"
              flex="~"
              justify="between"
              align="items-center"
              text="sm"
              border="rounded-2xl"
              outline="focus:none"
              >{{ t('user_details.form.roles') }} <bi-chevron-expand
            /></MenuButton>
            <MenuItems
              m="t-2"
              p="x-2 y-3"
              w="full"
              bg="gray-100 dark:dark-500"
              border="rounded-2xl"
              flex="~ col"
              align="items-start"
            >
              <MenuItem v-for="(role, index) in roles" :key="role.id" v-slot="{ active }">
                <button
                  :m="index !== 0 ? 't-2' : ''"
                  p="x-3 y-2"
                  w="full"
                  :text="`sm left ${active ? 'gray-700' : ''}`"
                  :bg="active ? 'neonGreen' : 'gray-200 dark:dark-400'"
                  border="rounded-xl"
                  outline="focus:outline-none"
                  @click="addRole(role.value)"
                >
                  {{ locale === 'en' ? role.name.en : role.name['zh-CN'] }}
                </button>
              </MenuItem>
            </MenuItems>
          </Menu>

          <input
            v-show="selectedRoles.find((r) => r.value === 'trainee')"
            name="trainee_id"
            v-model="trainee_id"
            :placeholder="t('user_details.form.trainee_id')"
            m="t-4"
            w="full"
            p="x-3 y-2"
            bg="gray-200 dark:dark-300"
            text="sm"
            border="rounded-2xl"
            outline="focus:none"
          />
          <span m="t-2" text="xs red-500">{{ traineeIdError }}</span>

          <button
            m="t-4"
            w="full"
            p="y-2"
            bg="neonGreen"
            text="sm gray-700"
            font="medium"
            border="rounded-2xl"
            outline="focus:none"
          >
            {{ t('user_details.submit') }}
          </button>
        </form>
        <div m="t-4" flex="~" align="items-center">
          <ThemeToggle />
          <LocaleToggle m="l-2" />
        </div>
      </template>

      <Spinner v-else animate="spin" m="t-8 x-auto" w="12" h="12" text="neonGreen" />
    </div>
  </Dialog>
</template>
