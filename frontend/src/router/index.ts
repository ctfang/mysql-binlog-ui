import { createRouter, createWebHashHistory } from 'vue-router'

export const AppMenu = [
  {
    path: '/',
    name: 'home',
    component: () => import('../views/HomeView.vue'),
    meta: {
      title: 'Home',
      icon: 'Connection'
    }
  },
  {
    path: '/binlogs',
    name: 'binlogs',
    component: () => import('../views/BinlogView.vue'),
    meta: {
      title: 'Logs',
      icon: 'plus'
    }
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: AppMenu
})

export default router
