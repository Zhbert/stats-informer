<template>
  <nav class="navbar navbar-expand-lg bg-light border_style">
    <div class="container-fluid">
      <router-link to="/" class="navbar-brand">Stats Informer</router-link>
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
            <router-link
              to="/statistics"
              class="nav-link active"
              aria-current="page"
              >Statistics</router-link
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
</template>

<script>
export default {
  name: "navbar-component",
};
</script>

<style scoped></style>
