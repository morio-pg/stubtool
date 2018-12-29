<template>
  <div class="authentication">
    <div v-if="userStatus" key="login" class="navbar-item">
      <p class="navbar-item">{{ user.displayName }}</p>
      <button type="button" class="button is-small is-info" @click="doLogout">
        Sign out
      </button>
    </div>
  </div>
</template>
<script>
import Firebase from "@/firebase";

export default {
  created: function() {
    Firebase.onAuth();
  },
  computed: {
    user() {
      return this.$store.getters.user;
    },
    userStatus() {
      return this.$store.getters.isSignedIn;
    }
  },
  methods: {
    doLogout() {
      Firebase.logout();
    }
  }
};
</script>
<style scoped>
.authentication {
  display: inline-block;
}
.navbar-item {
  color: #e5e5e5;
}
</style>
