import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/overview'
    },
    {
      path: "/overview",
      name: "概览",
      component: () => import('@/components/Overview/index.vue'),
    },
    {
      path: "/tool",
      name: "工具",
      component: () => import('@/components/Tool/index.vue'),
      children: [
        {
          path: "base64",
          component: import('@/components/Tool/sub/base64.vue')
        },
      ]
    },
    {
      path: "/news",
      name: "新闻",
      component: () => import('@/components/News/index.vue'),
    },
  ]
})

export default router
