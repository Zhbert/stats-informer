<template>
  <component :is="layout"><router-view /></component>
</template>

<script>
import MainLayout from "@/layouts/MainLayout.vue";
import GreetingsLayout from "@/layouts/GreetingsLayout.vue";
export default {
  computed: {
    layout() {
      return this.$route.meta.layout || "MainLayout";
    },
  },
  components: {
    MainLayout,
    GreetingsLayout,
  },
};
</script>

<style lang="scss">
@import "~bootstrap/dist/css/bootstrap.min.css";
</style>
