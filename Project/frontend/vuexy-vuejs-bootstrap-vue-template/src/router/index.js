import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  scrollBehavior() {
    return { x: 0, y: 0 }
  },
  routes: [
    {
      path: '/',
      name: 'deteksi',
      component: () => import('@/views/Deteksi.vue'),
      meta: {
        pageTitle: 'Deteksi',
        breadcrumb: [
          {
            text: 'Deteksi',
            active: true,
          },
        ],
      },
    },
    {
      path: '/result',
      name: 'result',
      component: () => import('@/views/Result.vue'),
      meta: {
        pageTitle: 'Result',
        breadcrumb: [
          {
            text: 'Result',
            active: true,
          },
        ],
      },
    },
    {
      path: '/instruction',
      name: 'instruction',
      component: () => import('@/views/Instruction.vue'),
      meta: {
        pageTitle: 'Instruction',
        breadcrumb: [
          {
            text: 'Instruction',
            active: true,
          },
        ],
      },
    },{
      path: '/list-deteksi',
      name: 'list-deteksi',
      component: () => import('@/views/ListDeteksi.vue'),
      meta: {
        pageTitle: 'List Deteksi',
        breadcrumb: [
          {
            text: 'List Deteksi',
            active: true,
          },
        ],
      },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/Login.vue'),
      meta: {
        layout: 'full',
      },
    },
    {
      path: '/error-404',
      name: 'error-404',
      component: () => import('@/views/error/Error404.vue'),
      meta: {
        layout: 'full',
      },
    },
    {
      path: '*',
      redirect: 'error-404',
    },
  ],
})

// ? For splash screen
// Remove afterEach hook if you are not using splash screen
router.afterEach(() => {
  // Remove initial loading
  const appLoading = document.getElementById('loading-bg')
  if (appLoading) {
    appLoading.style.display = 'none'
  }
})

export default router
