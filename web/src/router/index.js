import { createRouter, createWebHistory } from 'vue-router';


const DefaultLayout = () => import('@/layouts/DefaultLayout.vue');
const DeploymentList = () => import('@/views/DeploymentList.vue');


const routes = [
    {
        path: "/kubernetes",
        component: DefaultLayout,
        children: [
            {
            path: "deployments/list",
            name: "DeploymentList",
            component: DeploymentList,
            meta: { title: "DeploymentList" },
        },
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;