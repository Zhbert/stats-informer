import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";

const routes = [
  {
    path: "/",
    name: "main",
    meta: { layout: "GreetingsLayout" },
    component: HomeView,
  },
  {
    path: "/statistics",
    name: "stats",
    meta: { layout: "MainLayout" },
    component: () => import("@/views/StatisticsView.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
