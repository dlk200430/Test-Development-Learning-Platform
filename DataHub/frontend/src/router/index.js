import { createRouter, createWebHistory } from 'vue-router'
import Layout from '../components/Layout.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue')
  },
  {
    path: '/',
    component: Layout,
    children: [
      {
        path: '',
        redirect: '/dashboard'
      },
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../views/Dashboard.vue')
      },
      {
        path: 'testcases',
        name: 'TestCases',
        component: () => import('../views/TestCases.vue')
      },
      {
        path: 'testcases/create',
        name: 'CreateTestCase',
        component: () => import('../views/CreateTestCase.vue')
      },
      {
        path: 'testcases/:id/edit',
        name: 'EditTestCase',
        component: () => import('../views/EditTestCase.vue')
      },
      {
        path: 'reports',
        name: 'Reports',
        component: () => import('../views/Reports.vue')
      },
      {
        path: 'reports/:id',
        name: 'ReportDetail',
        component: () => import('../views/ReportDetail.vue')
      }
    ],
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  console.log('🛡️ 路由守卫触发')
  console.log('   ├─ 目标路径:', to.path)
  console.log('   ├─ 路由名称:', to.name)
  console.log('   ├─ 是否需要认证:', to.meta.requiresAuth)
  console.log('   ├─ Token存在:', !!token)
  
  if (to.path === '/login' || to.path === '/register') {
    if (token) {
      console.log('   └─ 已登录，重定向到 /dashboard')
      next('/dashboard')
    } else {
      console.log('   └─ 未登录，允许访问')
      next()
    }
  } else if (to.meta.requiresAuth && !token) {
    console.log('   └─ 需要认证但无Token，重定向到 /login')
    next('/login')
  } else {
    console.log('   └─ 验证通过，放行')
    next()
  }
})

export default router