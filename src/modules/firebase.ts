import firebase from 'firebase/app'
import 'firebase/auth'

const config = {
  apiKey: 'AIzaSyDP6wWLQAtklXfop3O7gDBiMBsd-86XXYw',
  authDomain: 'mentu-lxs.firebaseapp.com',
  projectId: 'mentu-lxs',
  storageBucket: 'mentu-lxs.appspot.com',
  messagingSenderId: '408026763213',
  appId: '1:408026763213:web:c138d2e999fcb364347e84',
}

firebase.initializeApp(config)

export const signIn = () => firebase.auth().signInWithPopup(new firebase.auth.GoogleAuthProvider())
export const signOut = () => firebase.auth().signOut()

export { firebase }
