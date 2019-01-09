import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    loading: true,
    user: {},
    status: false
  },
  mutations: {
    onSiteLoadCompleted(state) {
      state.loading = false;
    },
    onAuthStateChanged(state, user) {
      state.user = user;
    },
    onUserStatusChanged(state, status) {
      state.status = status;
    }
  },
  getters: {
    loading(state) {
      return state.loading;
    },
    user(state) {
      return state.user;
    },
    idToken(state) {
      return state.user._lat;
    },
    isSignedIn(state) {
      return state.status;
    }
  }
});
