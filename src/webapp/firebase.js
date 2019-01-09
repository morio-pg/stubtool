import firebase from "@firebase/app";
import "@firebase/auth";
import router from "./router";
import store from "./store";

var config = {
  apiKey: "AIzaSyDYWC_3ewDfIO8Oms9EUvCdaOuKDPSTHqw",
  authDomain: "stubtool.firebaseapp.com",
  databaseURL: "https://stubtool.firebaseio.com",
  projectId: "stubtool",
  storageBucket: "stubtool.appspot.com",
  messagingSenderId: "917561167215"
};

export default {
  init() {
    firebase.initializeApp(config);
    firebase.auth().setPersistence(firebase.auth.Auth.Persistence.LOCAL);
  },
  login() {
    const provider = new firebase.auth.TwitterAuthProvider();
    firebase.auth().signInWithRedirect(provider);
  },
  logout() {
    firebase.auth().signOut();
  },
  onAuth() {
    firebase.auth().onAuthStateChanged(user => {
      user = user ? user : {};
      var status = user.uid ? true : false;
      store.commit("onAuthStateChanged", user);
      store.commit("onUserStatusChanged", status);
      store.commit("onSiteLoadCompleted");
      if (!status && router.match(window.location.pathname).meta.requiresAuth) {
        router.push({ name: "Home" });
      }
    });
  }
};
