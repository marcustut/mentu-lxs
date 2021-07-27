<script setup lang="ts">
import { Dialog, DialogOverlay, Menu, MenuButton, MenuItems, MenuItem, TransitionRoot, TransitionChild } from '@headlessui/vue'
import { useAuth } from '@vueuse/firebase'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { firebase, signOut } from '~/modules/firebase'
import { isDark, toggleDark } from '~/logic'

const props = defineProps({
  class: String,
})

const drawerOpen = ref(false)
const { user } = useAuth(firebase.auth)
const { push } = useRouter()
const { locale } = useI18n()

const avatarUrl = user.value && user.value?.photoURL ? user.value.photoURL : `https://ui-avatars.com/api/?name=${user.value?.displayName?.replace(' ', '+')}`

const closeDrawer = () => drawerOpen.value = false
const openDrawer = () => drawerOpen.value = true

const toggleLocale = () => {
  if (locale.value === 'en')
    locale.value = 'zh-CN'
  else
    locale.value = 'en'
}

const routes = [
  { name: 'Home', icon: 'flat-color-icons:home', path: '/' },
  { name: 'Launching', icon: 'emojione-confetti-ball', path: '/launching' },
  { name: 'Linktree', icon: 'emojione-v1:evergreen-tree', path: '/linktree' },
]

const { t } = useI18n()
</script>

<template>
  <div flex="~" align="items-center" justify="between" :class="props.class">
    <button outline="focus:outline-none" @click="openDrawer">
      <heroicons-outline:menu-alt-2 w="8" h="8" />
    </button>
    <Menu>
      <MenuButton outline="focus:outline-none">
        <img w="10" h="10" object="cover" border="rounded-full" :src="avatarUrl" />
      </MenuButton>

      <MenuItems
        p="2"
        z="10"
        bg="gray-50 dark:dark-400"
        pos="absolute top-18 right-4"
        text="gray-700 dark:gray-200"
        border="rounded-2xl"
        shadow="xl"
      >
        <MenuItem v-slot="{ active }">
          <button
            :class="{ 'bg-neonGreen': active, 'text-textDark': active }"
            w="full"
            p="2"
            flex="~"
            font="medium"
            align="items-center"
            border="rounded-xl"
            outline="focus:outline-none"
            transition="colors duration-200 ease-in-out"
            @click="toggleDark"
          >
            <template v-if="isDark">
              <bx-bxs-moon v-if="isDark" m="r-2" />
              <span>{{ t('menu.toggle_theme_dark') }}</span>
            </template>
            <template v-else>
              <fa-solid:sun m="r-2" />
              <span>{{ t('menu.toggle_theme_light') }}</span>
            </template>
          </button>
        </MenuItem>
        <MenuItem v-slot="{ active }">
          <button
            :class="{ 'bg-neonGreen': active, 'text-textDark': active }"
            w="full"
            p="2"
            flex="~"
            font="medium"
            align="items-center"
            border="rounded-xl"
            outline="focus:outline-none"
            transition="colors duration-200 ease-in-out"
            @click="toggleLocale"
          >
            <uil-english-to-chinese m="r-2" />
            <span v-if="locale === 'en'">{{ t('menu.change_language_en') }}</span>
            <span v-else>{{ t('menu.change_language_cn') }}</span>
          </button>
        </MenuItem>
        <MenuItem v-slot="{ active }">
          <button
            :class="{ 'bg-neonGreen': active, 'text-textDark': active }"
            w="full"
            p="2"
            flex="~"
            font="medium"
            align="items-center"
            border="rounded-xl"
            outline="focus:outline-none"
            transition="colors duration-200 ease-in-out"
            @click="signOut"
          >
            <heroicons-outline:logout m="r-2" />
            <span>{{ t('menu.sign_out') }}</span>
          </button>
        </MenuItem>
      </MenuItems>
    </Menu>
  </div>

  <TransitionRoot :show="drawerOpen" as="template">
    <Dialog :open="drawerOpen" as="div" @close="closeDrawer">
      <DialogOverlay pos="fixed inset-0" backdrop="~ blur-[2px] brightness-75" />

      <TransitionChild
        as="template"
        enter="transform duration-300 ease-out"
        enter-from="-translate-x-full opacity-0"
        enter-to="translate-x-0 opacity-100"
        leave="transform duration-200 ease-in"
        leave-from="translate-x-0 opacity-100"
        leave-to="-translate-x-full opacity-0"
      >
        <div
          p="4"
          bg="gray-50 dark:dark-400"
          pos="fixed top-0 left-0"
          min-w="2/5"
          min-h="screen"
          text="gray-700 dark:gray-200"
        >
          <button
            v-for="(route, index) in routes"
            :key="route.name"
            :m="index !== 0 ? 't-2' : ''"
            w="full"
            p="x-4 y-2"
            bg="hover:neonGreen"
            text="hover:textDark"
            font="medium"
            flex="~"
            align="items-center"
            border="rounded-xl"
            outline="focus:outline-none"
            transition="colors duration-200 ease-in-out"
            @click="() => {
              push(route.path)
              closeDrawer()
            }"
          >
            <Icon :icon="route.icon" w="6" h="6" m="r-2" />
            {{ route.name }}
          </button>
        </div>
      </TransitionChild>
    </Dialog>
  </TransitionRoot>
</template>
