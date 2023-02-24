<template>
  <div class="container mt-2">
    <Navbar />
    <router-view />
  </div>
</template>

<script>
import Navbar from "@/components/app/navbar.vue";
export default {
  name: "MainLayout.vue",
  components: {
    Navbar,
  },
};
</script>

<style lang="css"></style>
