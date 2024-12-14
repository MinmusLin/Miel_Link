import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Layout from '@/layout'

export const constantRoutes = [
    {
        path: '/',
        redirect: '/home',
        hidden: true
    },
    {
        path: '/home',
        component: () => import('@/views/home/index'),
        meta: {
            title: '主页',
            icon: 'el-icon-s-home'
        }
    },
    {
        path: '/login',
        component: () => import('@/views/login/index'),
        hidden: true,
        meta: {
            title: '登录'
        }
    },
    {
        path: '/uplink',
        component: Layout,
        children: [{
            path: '',
            component: () => import('@/views/uplink/index'),
            meta: {
                title: '信息录入',
                icon: 'el-icon-edit-outline'
            }
        }]
    },
    {
        path: '/trace',
        component: Layout,
        children: [{
            path: '',
            component: () => import('@/views/trace/index'),
            meta: {
                title: '溯源查询',
                icon: 'el-icon-search'
            }
        }]
    },
    {
        path: 'hyperledger-explorer',
        component: Layout,
        children: [{
            path: 'http://43.156.142.179:8080',
            meta: {
                title: 'Hyperledger 浏览器',
                icon: 'el-icon-discover'
            }
        }]
    },
    {
        path: '*',
        redirect: '/home',
        hidden: true
    }
]

// noinspection JSCheckFunctionSignatures
const createRouter = () => new Router({
    mode: 'history',
    scrollBehavior: () => ({y: 0}),
    routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
    const newRouter = createRouter()
    // noinspection JSUnresolvedReference
    router.matcher = newRouter.matcher
}

export default router
