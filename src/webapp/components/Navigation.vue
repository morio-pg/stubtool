<template>
  <div id="nav">
    <div id="container">
      <nav class="navbar is-dark is-transparent">
        <div class="navbar-start">
          <h1 class="navbar-item title">Stub Tool</h1>
          <router-link class="navbar-item" :to="{ name: 'Home' }"
            >Home</router-link
          >
          <router-link
            class="navbar-item"
            :to="{ name: 'Create' }"
            v-show="isSignedIn"
            >Create</router-link
          >
          <router-link
            class="navbar-item"
            :to="{ name: 'Account' }"
            v-show="isSignedIn"
            >Account</router-link
          >
        </div>
        <div class="navbar-end"><Authentication /></div>
      </nav>
    </div>
  </div>
</template>
<script>
import Authentication from "@/components/Authentication.vue";

export default {
  components: {
    Authentication
  },
  computed: {
    isSignedIn() {
      return this.$store.getters.isSignedIn;
    }
  }
};
</script>
<style scoped>
#nav {
  background-color: #363636;
}
#container {
  max-width: 960px;
  width: 960px;
  margin: auto;
  position: relative;
  padding: 0rem;
}
#nav .navbar-item {
  color: #e5e5e5;
}
#nav .title {
  color: #ffffff;
  height: 56px;
  margin: auto;
  padding: 0rem;
  padding-right: 1.5rem;
}
#nav a {
  font-weight: bold;
}
#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
