import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/uplink',
    children: [{
      path: 'uplink',
      name: 'Uplink',
      component: () => import('@/views/uplink/index'),
      meta: { title: '溯源信息录入', icon: 'el-icon-edit-outline' }
    }]
  },
  {
    path: '/trace',
    component: Layout,
    children: [
      {
        path: '',
        name: 'Trace',
        component: () => import('@/views/trace/index.vue'),
        meta: { title: '溯源查询', icon: 'el-icon-search' }
      },
      {
        path: ':traceability_code',
        name: 'TraceDetail',
        component: () => import('@/views/trace/index.vue'),
        meta: { title: '溯源查询详情' },
        hidden: true
      }
    ]
  },
  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'http://119.45.247.29:8080',
        meta: { title: '区块链浏览器', icon: 'el-icon-discover' }
      }
    ]
  },
  {
    path: '/build',
    component: Layout,
    redirect: '/build',
    children: [{
      path: 'build',
      name: 'Build',
      component: () => import('@/views/build/index'),
      meta: { title: '构建任意溯源系统', icon: 'el-icon-s-tools' }
    }]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
