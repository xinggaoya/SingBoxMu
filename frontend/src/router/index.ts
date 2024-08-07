import {createRouter, createWebHashHistory} from 'vue-router'
import Layout from '@/components/layout/Layout.vue'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'index',
      component: Layout,
      redirect: '/chat',
      children: [
        {
          path: '/chat/:id?',
          name: 'chat',
          component: () => import('@/views/HomeView.vue')
        },
      ]
    },
    {
      path: '/login',
      name: 'Login',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('@/views/ChatLogin.vue')
    }
  ]
})

export default router
