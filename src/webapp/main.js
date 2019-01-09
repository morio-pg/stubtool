import Vue from "vue";
import App from "./App.vue";
import Buefy from "buefy";
import "buefy/dist/buefy.min.css";
import "@mdi/font/css/materialdesignicons.min.css";
import Firebase from "./firebase";
import VeeValidate from "vee-validate";
import ja from "vee-validate/dist/locale/ja";
import router from "./router";
import store from "./store";
import * as filters from "./filters";

Firebase.init();
Vue.use(Buefy);
Vue.use(VeeValidate, {
  locale: "ja",
  dictionary: {
    ja: ja
  }
});

router.beforeEach((to, from, next) => {
  if (
    to.matched.some(record => record.meta.requiresAuth) &&
    !store.getters.loading &&
    !store.getters.isSignedIn
  ) {
    next({ name: "Home" });
  } else {
    next();
  }
});

Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key]);
});

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
