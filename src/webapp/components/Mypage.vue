<template>
  <div>
    <b-loading :active.sync="isLoading"></b-loading>

    <nav class="breadcrumb" aria-label="breadcrumbs">
      <ul>
        <li class="is-active"><a href="#" aria-current="page">Home</a></li>
      </ul>
    </nav>

    <b-table :data="data" :mobile-cards="false" hoverable>
      <template slot-scope="props">
        <b-table-column label="ファイル名" width="205">
          {{ props.row.filename | truncate(19) }}
        </b-table-column>

        <b-table-column label="MIMEタイプ" width="140">
          {{ props.row.mimeType }}
        </b-table-column>

        <b-table-column label="概要">
          {{ props.row.description | truncate(19) }}
        </b-table-column>

        <b-table-column label="更新日時" width="150" centered>
          {{ props.row.updatedAt | dateFormat }}
        </b-table-column>

        <b-table-column label="操作" width="70" centered>
          <router-link
            tag="button"
            class="button is-small"
            :to="{ name: 'Detail', params: { stubId: props.row.stubId } }"
            >Detail</router-link
          >
        </b-table-column>
      </template>

      <template slot="empty">
        <section class="section" v-show="!isLoading">
          <div class="content has-text-grey has-text-centered">
            <p><b-icon icon="emoticon-sad" size="is-large"></b-icon></p>
            <p>Nothing here.</p>
          </div>
        </section>
      </template>
    </b-table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      no: 0,
      data: [],
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
        .get("/api/admin", {
          headers: {
            Authorization: `Bearer ${this.$store.getters.idToken}`
          }
        })
        .then(
          function(response) {
            this.data = response.data;
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
          }.bind(this)
        );
    }
  }
};
</script>
