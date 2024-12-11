import { createRouter, createWebHistory } from 'vue-router';


const DeploymentList = () => import('@/views/DeploymentList.vue');

const route = [
    {
        path: "/deployments",
        name: "Deployments",
        component: DeploymentList
    }
]

const route = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;