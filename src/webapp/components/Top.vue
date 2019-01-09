<template>
  <div>
    <article class="message is-small">
      <div class="message-body">
        <pre>
$ curl -i https://stubtool.herokuapp.com/api/stubs/bgqvia8igc736454afpg/getHelloWorld
HTTP/1.1 200 OK
Server: Cowboy
Connection: keep-alive
Content-Type: application/json
Date: Wed, 09 Jan 2019 14:45:48 GMT
Content-Length: 27
Via: 1.1 vegur

{"message": "Hello World!"}
</pre
        >
      </div>
    </article>

    <div class="content has-text-centered">
      <div class="description is-size-7">
        <p>
          <strong>Stub Tool</strong
          >ではJSONやXMLのスタブを作成することができます
        </p>
        <p>その他にもステータスコードや待機時間の設定を行うことができます</p>
        <p>このサイトはβ版になります</p>
        <p class="has-text-danger">
          大事な情報などは登録しないようにご注意ください
        </p>
      </div>

      <button type="button" class="button twitter-button" @click="doLogin">
        <span class="twitter-button__icon">
          <img src="@/assets/twitter.svg" />
        </span>
        <span class="twitter-button__text">Sign in with Twitter</span>
      </button>
    </div>
  </div>
</template>
<script>
import Firebase from "@/firebase";

export default {
  methods: {
    doLogin() {
      Firebase.login();
    }
  }
};
</script>
<style scoped>
.twitter-button {
  height: 40px;
  border-width: 0;
  background: #55acee;
  color: #ffffff;
  border-radius: 5px;
  white-space: nowrap;
  box-shadow: 1px 1px 0px 1px rgba(0, 0, 0, 0.05);
  transition-property: background-color, box-shadow;
  transition-duration: 150ms;
  transition-timing-function: ease-in-out;
  padding: 0;
}
.twitter-button:focus,
.twitter-button:hover {
  box-shadow: 1px 4px 5px 1px rgba(0, 0, 0, 0.1);
}
.twitter-button:active {
  box-shadow: none;
  transition-duration: 10ms;
}
.twitter-button__icon {
  display: inline-block;
  vertical-align: middle;
  margin: 8px -12px 8px 20px;
  width: 18px;
  height: 18px;
  box-sizing: border-box;
}
.twitter-button__text {
  display: inline-block;
  vertical-align: middle;
  padding: 0 20px;
  font-size: 14px;
  font-weight: bold;
  font-family: "Roboto", arial, sans-serif;
}
.message {
  margin-top: 1.5rem;
}
.description {
  margin: 2.5rem;
}
</style>
