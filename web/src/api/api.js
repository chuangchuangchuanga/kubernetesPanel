import request from "@/utils/request.js";
import service from "@/utils/request.js";


export function getNamespaceList() {
    return service({
        url: "/namespace",
        method: "GET",
    })
}

export function getDeploymentList(data){
    return service({
        url: "/deploymentlist",
        method: "POST",
        data: data,
        headers: {
            'Content-Type': 'application/json',
        }
    })
}


export function postRestartDeployment(data){
    return service({
        url: "/deploymentrestart",
        method: "POST",
        data: data,
        headers: {
            'Content-Type': 'application/json',
        }
    })
}

export function getDeploymentPodList(data){
    return service({
        url: "/getdeploymentpods",
        method: "POST",
        data: data,
        headers: {
            'Content-Type': 'application/json',
        }
    })
}


