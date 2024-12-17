import { createRouter, createWebHistory } from 'vue-router';


const DefaultLayout = () => import('@/layouts/DefaultLayout.vue');
const DeploymentList = () => import('@/views/DeploymentList.vue');
const podLogPage = () => import('@/views/podLogPage.vue');


const routes = [
    {
        path: "/kubernetes",
        component: DefaultLayout,
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
    }

]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;