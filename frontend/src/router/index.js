import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "main",
    meta: { layout: "GreetingsLayout" },
  },
  {
    path: "/statistics",
    name: "stats",
    meta: { layout: "MainLayout" },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
