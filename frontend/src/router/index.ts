import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/share/:token',
      name: 'share',

      component: () => import('../views/ShareView.vue')
    },
    {
      path: '/check/:token/:token1/:token2',
      name: 'check',

      component: () => import('../views/DetailsView.vue')
    },
    {
       path: '/:pathMatch(.*)*',
       name: 'notfound',
      component: () => import('../views/NotFoundView.vue')

    }

  ]
})

export default router
