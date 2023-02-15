<template>
  <div class="container mt-2">
    <nav class="navbar navbar-expand-lg bg-light border_style">
      <div class="container-fluid">
        <a class="navbar-brand" href="/">Stats Informer</a>
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNavDropdown"
          aria-controls="navbarNavDropdown"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNavDropdown">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <a
                class="nav-link active"
                aria-current="page"
                href="/github-short"
                >GitHub</a
              >
            </li>
            <li class="nav-item dropdown">
              <a
                class="nav-link dropdown-toggle"
                href="#"
                role="button"
                data-bs-toggle="dropdown"
                aria-expanded="false"
              >
                Help
              </a>
              <ul class="dropdown-menu">
                <li><a class="dropdown-item" href="#">About</a></li>
              </ul>
            </li>
          </ul>
          <div class="credits">
            v.0.3 | Made with
            <span class="page__icon page__icon_heart"></span> by
            <a href="https://flant.com" target="_blank">Flant</a>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
export default {
  name: "MainLayout.vue",
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
