<template>
  <div class="container mt-2">
    <Navbar />
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

.border_style {
  border: solid 1px black;
  border-radius: 5px;
}
</style>
