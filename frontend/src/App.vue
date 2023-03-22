<template>
  <component :is="layout"><router-view /></component>
</template>

<script>
import MainLayout from "@/layouts/MainLayout.vue";
import GreetingsLayout from "@/layouts/GreetingsLayout.vue";
import axios from "axios";
export default {
  computed: {
    layout() {
      return this.$route.meta.layout || "MainLayout";
    },
  },
  data() {
    return {
      version: "-",
    };
  },
  components: {
    MainLayout,
    GreetingsLayout,
  },
  created() {
    axios.get("/api/v1/get-version").then((response) => {
      this.version = response.data.version;
      console.log(this.version);
    });
  },
};
</script>

<style lang="scss">
@import "~bootstrap/dist/css/bootstrap.min.css";
</style>
