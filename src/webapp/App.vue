<template>
  <div id="app">
    <div class="is-hidden-desktop">
      <strong>This page can only be displayed on PC.</strong>
    </div>
    <div class="is-hidden-touch">
      <b-loading :active.sync="loading"></b-loading>
      <div id="main" v-show="!loading">
        <Navigation />
        <router-view />
        <Bottom />
      </div>
    </div>
  </div>
</template>
<script>
import Navigation from "@/components/Navigation.vue";
import Bottom from "@/components/Bottom.vue";

export default {
  components: {
    Navigation,
    Bottom
  },
  computed: {
    loading() {
      return this.$store.getters.loading;
    }
  }
};
</script>
<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
#main {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
#container {
  max-width: 960px;
  width: 960px;
  margin: auto;
  margin-top: 0rem;
  position: relative;
  padding: 1.5rem;
}
.breadcrumb {
  height: 27px;
}
</style>
