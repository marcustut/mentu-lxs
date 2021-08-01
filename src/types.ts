import { ViteSSGContext } from 'vite-ssg';
import { firebase } from '~/modules/firebase';

export type UserModule = (ctx: ViteSSGContext) => void;

export type Notification = {
  id: string;
  icon: string;
  link: string;
  title: { en: string; 'zh-CN': string };
  description: { en: string; 'zh-CN': string };
  active: boolean;
};

export type Event = {
  id: string;
  icon: string;
  start_datetime: firebase.firestore.Timestamp;
  title_link: string;
  description_link: string;
  title: { en: string; 'zh-CN': string };
  description: { en: string; 'zh-CN': string };
};

export type Launching = {
  count: number;
  participants: number;
  available: boolean;
  start_datetime: firebase.firestore.Timestamp;
  end_datetime: firebase.firestore.Timestamp;
};

export type Message = {
  content: string;
  sentBy: { name: string; avatar_url: string };
  createdAt: firebase.firestore.Timestamp;
};

export type User = {
  uid: string;
  trainee_id: string | null;
  displayName: string | null;
  email: string | null;
  name: { en: string | null; 'zh-CN': string | null };
  roles: string[];
  age: number | null;
  phone_number: string | null;
};
