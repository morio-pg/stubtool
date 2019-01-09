<template>
  <div id="container">
    <b-loading :active.sync="isLoading"></b-loading>

    <nav class="level">
      <div class="level-left">
        <div class="level-item">
          <nav class="breadcrumb" aria-label="breadcrumbs">
            <ul>
              <li>
                <router-link class="navbar-item" :to="{ name: 'Home' }"
                  >Home</router-link
                >
              </li>
              <li class="is-active">
                <a href="#" aria-current="page">Detail</a>
              </li>
            </ul>
          </nav>
        </div>
      </div>

      <div class="level-right">
        <div class="level-item">
          <router-link
            tag="button"
            class="button is-small"
            :to="{
              name: 'Edit',
              params: { stubId: this.$route.params.stubId }
            }"
            >Edit</router-link
          >
        </div>
        <div class="level-item">
          <button
            type="button"
            class="button is-small is-danger"
            @click="confirmDelete"
          >
            Delete
          </button>
        </div>
      </div>
    </nav>

    <div class="box">
      <b-field horizontal label="URL"
        ><a :href="url" target="_blank">{{ url }}</a></b-field
      >
      <b-field horizontal label="更新日時"
        ><div v-show="updatedAt">{{ updatedAt | dateFormat }}</div></b-field
      >
      <b-field horizontal label="ファイル名">{{ filename }}</b-field>
      <b-field horizontal label="概要">{{ description }}</b-field>
      <b-field horizontal label="ステータスコード">{{ statusCode }}</b-field>
      <b-field horizontal label="待機時間（ms）">{{ wait }}</b-field>
      <b-field horizontal label="MIMEタイプ">{{ mimeType }}</b-field>
      <b-field horizontal label="内容">
        <b-input type="textarea" v-model="content" rows="8" readonly></b-input>
      </b-field>
    </div>
  </div>
</template>
<script>
import axios from "axios";

export default {
  data() {
    return {
      url: null,
      updatedAt: null,
      filename: null,
      description: null,
      statusCode: null,
      wait: null,
      mimeType: null,
      content: null,
      isLoading: true
    };
  },
  computed: {
    loading() {
      return this.$store.getters.loading;
    }
  },
  created: function() {
    if (!this.$store.getters.loading) {
      this.getData();
    }
  },
  watch: {
    loading() {
      this.getData();
    }
  },
  methods: {
    getData() {
      if (!this.$store.getters.isSignedIn) {
        this.isLoading = false;
        return;
      }

      axios
        .get("/api/admin/" + this.$route.params.stubId, {
          headers: {
            Authorization: `Bearer ${this.$store.getters.idToken}`
          }
        })
        .then(
          function(response) {
            this.url =
              window.location.protocol +
              "//" +
              window.location.host +
              "/api/stubs/" +
              this.$route.params.stubId +
              "/" +
              response.data.filename;
            this.updatedAt = response.data.updatedAt;
            this.filename = response.data.filename;
            this.description = response.data.description
              ? response.data.description
              : "-";
            this.statusCode = response.data.statusCode;
            this.wait = response.data.wait + " ms";
            this.mimeType = response.data.mimeType;
            this.content = response.data.content;
            this.isLoading = false;
          }.bind(this)
        )
        .catch(
          function() {
            this.isLoading = false;
            this.$toast.open({
              message: "スタブの取得に失敗しました",
              type: "is-danger"
            });
            this.$router.push({ name: "Home" });
          }.bind(this)
        );
    },
    confirmDelete() {
      this.$dialog.confirm({
        message: "スタブを削除してよろしいですか？",
        type: "is-danger",
        onConfirm: () => {
          this.isLoading = true;

          axios
            .delete("/api/admin/" + this.$route.params.stubId, {
              headers: {
                Authorization: `Bearer ${this.$store.getters.idToken}`
              }
            })
            .then(
              function() {
                this.isLoading = false;
                this.$toast.open({
                  message: "スタブを削除しました",
                  type: "is-success"
                });
                this.$router.push({ name: "Home" });
              }.bind(this)
            )
            .catch(
              function() {
                this.isLoading = false;
                this.$toast.open({
                  message: "スタブの削除に失敗しました",
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
