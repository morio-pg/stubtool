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
      <h3>ページが見つかりませんでした</h3>
    </div>
  </div>
</template>
<style scoped>
.content {
  padding: 3rem;
}
</style>
