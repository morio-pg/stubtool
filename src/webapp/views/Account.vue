<template>
  <div id="container">
    <b-loading :active.sync="isLoading"></b-loading>

    <nav class="breadcrumb" aria-label="breadcrumbs">
      <ul>
        <li>
          <router-link class="navbar-item" :to="{ name: 'Home' }"
            >Home</router-link
          >
        </li>
        <li class="is-active"><a href="#" aria-current="page">Account</a></li>
      </ul>
    </nav>

    <div class="content">
      <p>
        <i class="mdi mdi-delete-outline mdi-22px"></i
        ><strong>アカウント削除</strong>
      </p>
      <p>作成したスタブも全て削除されます。ご注意ください。</p>
      <p>
        <button type="button" class="button is-danger" @click="confirmDelete">
          アカウントを削除する
        </button>
      </p>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import firebase from "@firebase/app";

export default {
  data() {
    return {
      isLoading: false
    };
  },
  methods: {
    confirmDelete() {
      this.$dialog.confirm({
        message: "アカウントを削除してよろしいですか？",
        type: "is-danger",
        onConfirm: () => {
          this.isLoading = true;

          axios
            .delete("/api/account", {
              headers: {
                Authorization: `Bearer ${this.$store.getters.idToken}`
              }
            })
            .then(
              function() {
                this.isLoading = false;
                this.$toast.open({
                  message: "アカウントを削除しました",
                  type: "is-success"
                });
                firebase.auth().signOut();
                this.$router.push({ name: "Home" });
              }.bind(this)
            )
            .catch(
              function() {
                this.isLoading = false;
                this.$toast.open({
                  message: "アカウントの削除に失敗しました",
                  type: "is-danger"
                });
              }.bind(this)
            );
        }
      });
    }
  }
};
</script>
<style scoped>
.content {
  width: 480px;
  margin: auto;
  padding: 1.5rem;
}
</style>
