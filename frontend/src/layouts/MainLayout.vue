<template>
  <div class="container mt-2">
    <Navbar />
    <div class="container mt-3">
      <div class="row">
        <div class="col-3 border_style me-1">5</div>
        <div class="col border_style ms-1">6</div>
      </div>
    </div>
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

<style lang="css">
.border_style {
  border: solid 1px black;
  border-radius: 5px;
}
</style>
