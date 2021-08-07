import firebase from 'firebase/app'
import 'firebase/auth'
import 'firebase/firestore'
import 'firebase/storage'

const config = {
  apiKey: import.meta.env.VITE_FIREBASE_API_KEY,
  authDomain: import.meta.env.VITE_FIREBASE_AUTH_DOMAIN,
  projectId: import.meta.env.VITE_FIREBASE_PROJECT_ID,
  storageBucket: import.meta.env.VITE_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: import.meta.env.VITE_FIREBASE_MESSAGING_SENDER_ID,
  appId: import.meta.env.VITE_FIREBASE_APP_ID,
}

firebase.initializeApp(config)

export const signIn = async(): Promise<{ user: firebase.auth.UserCredential | null; error: Error | null }> => {
  try {
    const user = await firebase.auth().signInWithPopup(new firebase.auth.GoogleAuthProvider())
    return { user, error: null }
  }
  catch (error) {
    return { user: null, error }
  }
}
export const signOut = () => firebase.auth().signOut()

export const db = firebase.firestore()

export const storage = firebase.storage()

export { firebase }
