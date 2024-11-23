import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue')
    },
    {
      path: '/data',
      name: 'data',
      component: () => import('../views/DataView.vue')
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('../views/LogoutView.vue')
    },
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  if (to.name === 'login' && authStore.isLoggedIn) {
    next({ name: 'data' })
  } else if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    next({ name: 'login' })
  } else {
    next()
  }
})

