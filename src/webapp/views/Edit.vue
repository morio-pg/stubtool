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
        <li>
          <router-link
            class="navbar-item"
            :to="{
              name: 'Detail',
              params: { stubId: this.$route.params.stubId }
            }"
            >Detail</router-link
          >
        </li>
        <li class="is-active"><a href="#" aria-current="page">Edit</a></li>
      </ul>
    </nav>

    <div class="box">
      <form @submit.prevent="validateBeforeSubmit">
        <b-field
          horizontal
          label="ファイル名"
          :type="{ 'is-danger': errors.has('filename') }"
          :message="errors.first('filename')"
        >
          <b-input
            name="filename"
            data-vv-as="ファイル名"
            v-model="filename"
            v-validate="'required|alpha_dash|min:3'"
            placeholder="Filename"
            maxlength="30"
            expanded
          ></b-input>
        </b-field>

        <b-field horizontal label="概要">
          <b-input
            name="description"
            v-model="description"
            placeholder="Description"
            maxlength="40"
            expanded
          ></b-input>
        </b-field>

        <b-field
          horizontal
          label="ステータスコード"
          :type="{ 'is-danger': errors.has('statusCode') }"
          :message="errors.first('statusCode')"
        >
          <b-input
            name="statusCode"
            data-vv-as="ステータスコード"
            v-model="statusCode"
            v-validate="'required|numeric|between:100,599'"
            placeholder="StatusCode"
            maxlength="3"
          ></b-input>
        </b-field>

        <b-field
          horizontal
          label="待機時間（ms）"
          :type="{ 'is-danger': errors.has('wait') }"
          :message="errors.first('wait')"
        >
          <b-input
            name="wait"
            data-vv-as="待機時間"
            v-model="wait"
            v-validate="'required|numeric|max_value:15000'"
            placeholder="Wait"
            maxlength="5"
          ></b-input>
        </b-field>

        <b-field
          horizontal
          label="MIMEタイプ"
          :type="{ 'is-danger': errors.has('mimeTypeKey') }"
          :message="errors.first('mimeTypeKey')"
        >
          <b-select
            name="mimeTypeKey"
            data-vv-as="MIMEタイプ"
            v-model="mimeTypeKey"
            v-validate="'required|included:1,2'"
            placeholder="MimeType"
          >
            <option value="1">application/json</option>
            <option value="2">application/xml</option>
          </b-select>
        </b-field>

        <b-field
          horizontal
          label="内容"
          :type="{ 'is-danger': errors.has('content') }"
          :message="errors.first('content')"
        >
          <b-input
            type="textarea"
            name="content"
            data-vv-as="内容"
            v-model="content"
            v-validate="'required'"
            placeholder="Content"
            rows="8"
            maxlength="5000"
          ></b-input>
        </b-field>

        <b-field horizontal
          ><!-- Label left empty for spacing -->
          <div class="field is-grouped">
            <p class="control">
              <button type="submit" class="button is-success">更新する</button>
            </p>
            <p class="control">
              <router-link
                class="button is-light"
                :to="{
                  name: 'Detail',
                  params: { stubId: this.$route.params.stubId }
                }"
                >キャンセル</router-link
              >
            </p>
          </div>
        </b-field>
      </form>
    </div>
  </div>
</template>
<script>
import axios from "axios";

export default {
  data() {
    return {
      filename: null,
      description: null,
      statusCode: null,
      wait: null,
      mimeTypeKey: null,
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
            this.filename = response.data.filename;
            this.description = response.data.description;
            this.statusCode = response.data.statusCode;
            this.wait = response.data.wait;
            this.mimeTypeKey = response.data.mimeTypeKey;
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
            this.$router.push({
              name: "Detail",
              params: { stubId: this.$route.params.stubId }
            });
          }.bind(this)
        );
    },
    validateBeforeSubmit() {
      this.$validator.validateAll().then(result => {
        if (result) {
          this.isLoading = true;

          let params = new URLSearchParams();
          params.append("filename", this.filename);
          params.append("description", this.description);
          params.append("statusCode", this.statusCode);
          params.append("wait", this.wait);
          params.append("mimeTypeKey", this.mimeTypeKey);
          params.append("content", this.content);

          axios
            .put("/api/admin/" + this.$route.params.stubId, params, {
              headers: {
                Authorization: `Bearer ${this.$store.getters.idToken}`
              }
            })
            .then(
              function() {
                this.isLoading = false;
                this.$toast.open({
                  message: "スタブを更新しました",
                  type: "is-success"
                });
                this.$router.push({
                  name: "Detail",
                  params: { stubId: this.$route.params.stubId }
                });
              }.bind(this)
            )
            .catch(
              function() {
                this.isLoading = false;
                this.$toast.open({
                  message: "スタブの更新に失敗しました",
                  type: "is-danger"
                });
              }.bind(this)
            );
        } else {
          this.$toast.open({
            message: "入力内容に不備があります",
            type: "is-danger"
          });
        }
      });
    }
  }
};
</script>
