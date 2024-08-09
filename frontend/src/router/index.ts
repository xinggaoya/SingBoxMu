import {createRouter, createWebHashHistory} from 'vue-router'
import Layout from '@/components/layout/Layout.vue'

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'index',
            component: Layout,
            children: [
                {
                    path: '/',
                    name: 'Home',
                    component: () => import('@/views/home/HomeView.vue')
                },
                {
                    path: '/sub',
                    name: 'SubView',
                    component: () => import('@/views/sub/SubView.vue')
                },
                {
                    path: '/log',
                    name: 'Log',
                    component: () => import('@/views/log/LogView.vue')
                },
                {
                    path: '/setting',
                    name: 'Setting',
                    component: () => import('@/views/setting/SettingView.vue')
                },
            ]
        },

    ]
})

export default router
