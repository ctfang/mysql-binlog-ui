import { createRouter, createWebHashHistory } from "vue-router";
import About from "@/component/icons/About.vue";
import BinlogIcon from "@/component/icons/Binlog.vue";
import AddIcon from "@/component/icons/Add.vue";

export const AppMenu = [
  {
    path: "/",
    name: "home",
    component: () => import("../views/HomeView.vue"),
    meta: {
      title: "Home",
      icon: BinlogIcon,
    },
  },
  {
    path: "/binlogs",
    name: "binlogs",
    component: () => import("../views/BinlogView.vue"),
    meta: {
      title: "Logs",
      icon: AddIcon,
    },
  },
  {
    path: "/about",
    name: "about",
    component: () => import("../views/AboutView.vue"),
    meta: {
      title: "About",
      icon: About,
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: AppMenu,
});

export default router;
