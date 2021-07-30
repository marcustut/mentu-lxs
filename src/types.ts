import { ViteSSGContext } from 'vite-ssg'
import { firebase } from '~/modules/firebase'

export type UserModule = (ctx: ViteSSGContext) => void

export type Notification = {
  id: string
  icon: string
  link: string
  title: { en: string; 'zh-CN': string }
  description: { en: string; 'zh-CN': string }
}

export type Event = {
  id: string
  icon: string
  updatedAt: firebase.firestore.Timestamp
  title_link: string
  description_link: string
  title: { en: string; 'zh-CN': string }
  description: { en: string; 'zh-CN': string }
}
