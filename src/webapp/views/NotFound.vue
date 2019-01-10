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
        <li class="is-active"><a href="#" aria-current="page">NotFound</a></li>
      </ul>
    </nav>

    <div class="content has-text-centered">
      <h3>指定されたページが見つかりません</h3>
      <p>URLに間違いがないかご確認ください</p>
      <p>ページの再読み込みをお試しください</p>
    </div>
  </div>
</template>
<style scoped>
.content {
  padding: 3rem;
}
</style>
