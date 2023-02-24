<template>
  <div class="px-4 py-5 my-5 text-center">
    <img
      src="@/assets/ant.png"
      style="width: 35%"
      class="mb-2"
      alt="Flant ant with papers"
    />
    <h1 class="display-5 fw-bold">Welcome to Stats Informer</h1>
    <div class="col-lg-6 mx-auto">
      <p class="lead mb-4">
        Here you can track changes in the number of stars in the GitHub
        repositories that you have subscribed to in the service
      </p>
      <div class="d-grid gap-2 d-sm-flex justify-content-sm-center">
        <router-link
          to="/statistics"
          class="btn btn-outline-secondary btn-lg px-4"
          >Start Tracking</router-link
        >
      </div>
    </div>
    <div class="credits mt-3">
      v.{{ version }} | Made with
      <span class="page__icon page__icon_heart"></span> by
      <a href="https://flant.com" target="_blank">Flant</a>
    </div>
  </div>
</template>

<script>
export default {
  name: "HomeView",
  data() {
    return {
      version: this.$root.$data.version,
    };
  },
};
</script>

<style lang="css">
.credits {
  font-size: 14px;
  font-weight: normal;
  font-style: normal;
  font-stretch: normal;
  line-height: 1.71;
  letter-spacing: -0.3px;
  color: #00122c;
}
.page__icon {
  display: inline-block;
  background-position: center center;
  background-repeat: no-repeat;
  background-size: contain;
}

.page__icon_heart {
  background-image: url("@/assets/heart.png");
  width: 12px;
  height: 12px;
}
</style>
