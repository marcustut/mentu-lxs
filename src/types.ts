import { ViteSSGContext } from 'vite-ssg';
import { firebase } from '~/modules/firebase';

type PrependNextNum<A extends Array<unknown>> = A['length'] extends infer T
  ? ((t: T, ...a: A) => void) extends (...x: infer X) => void
    ? X
    : never
  : never;

type EnumerateInternal<A extends Array<unknown>, N extends number> = {
  0: A;
  1: EnumerateInternal<PrependNextNum<A>, N>;
}[N extends A['length'] ? 0 : 1];

export type Enumerate<N extends number> = EnumerateInternal<[], N> extends (infer E)[] ? E : never;

export type Range<FROM extends number, TO extends number> = Exclude<Enumerate<TO>, Enumerate<FROM>>;

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
  avatar_url?: string;
  trainee_id: string | null;
  displayName: string | null;
  email: string | null;
  name: { en: string | null; 'zh-CN': string | null };
  roles: string[];
  age: number | null;
  phone_number: string | null;
};

export type Disc = {
  uid: string;
  displayName: string | null;
  responses: Record<Range<0, 40>, 'a' | 'b' | 'c' | 'd'>;
};

export type Gift = {
  uid: string;
  displayName: string | null;
  responses: Record<number, 0 | 1 | 2 | 3 | 4>;
};

export type Report = {
  uid: string;
  name: string;
  scores: {
    time: number;
    communication: number;
    willingness: number;
    responsibility: number;
    maturity: number;
  };
  strength: string;
  words: string;
  total: number;
  grand_total: number;
};

export type Group = {
  id: string;
  name: string;
  video: {
    caption: string;
    likes: number;
    name: string;
    src: string;
    postedAt: firebase.firestore.Timestamp;
    tags: string[];
    theme: string;
  };
};

export type InboxMessage = {
  attachment: string[];
  created_at: firebase.firestore.Timestamp;
  message: string;
  receiver: string;
  sender: string;
};
