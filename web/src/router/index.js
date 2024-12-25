import { createRouter, createWebHistory } from 'vue-router';


const DefaultLayout = () => import('@/layouts/DefaultLayout.vue');
const DeploymentList = () => import('@/views/DeploymentList.vue');
const podLogPage = () => import('@/views/podLogPage.vue');


const routes = [
    {
        path: "/kubernetes",
        name: "Kubernetes",
        component: DefaultLayout,
        meta: {title: 'Kubernetes'},
        children: [
            {
            path: "deployments/list",
            name: "DeploymentLists",
            component: DeploymentList,
            meta: { title: "DeploymentList" },
        },
        ]
    },
    {
        path: "/getpodlogs",
        name: "podLogPage",
        component: podLogPage,
        meta: { title: "PodLogPage" },
    }

]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    // 检查目标路由的 meta 对象是否有 title 字段
    if (to.meta && to.meta.title) {
        // 设置文档标题为路由的 meta.title
        document.title = to.meta.title;
    } else {
        // 设置一个默认标题
        document.title = '默认标题';
    }
    next();
});


export default router;