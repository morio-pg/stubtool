<template>
  <div id="container">
    <Mypage v-show="isSignedIn" />
    <Top v-show="!isSignedIn" />
  </div>
</template>
<script>
import Mypage from "@/components/Mypage.vue";
import Top from "@/components/Top.vue";

export default {
  components: {
    Mypage,
    Top
  },
  computed: {
    isSignedIn() {
      return this.$store.getters.isSignedIn;
    }
  }
};
</script>
