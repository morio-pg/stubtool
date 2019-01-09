import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      name: "Home",
      path: "/",
      component: () => import("./views/Home.vue")
    },
    {
      name: "Create",
      path: "/create",
      component: () => import("./views/Create.vue"),
      meta: { requiresAuth: true }
    },
    {
      name: "Detail",
      path: "/stubs/:stubId",
      component: () => import("./views/Detail.vue"),
      meta: { requiresAuth: true }
    },
    {
      name: "Edit",
      path: "/stubs/:stubId/edit",
      component: () => import("./views/Edit.vue"),
      meta: { requiresAuth: true }
    },
    {
      name: "Account",
      path: "/account",
      component: () => import("./views/Account.vue"),
      meta: { requiresAuth: true }
    },
    {
      path: "*",
      component: () => import("./views/NotFound.vue")
    }
  ],
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { x: 0, y: 0 };
    }
  }
});
