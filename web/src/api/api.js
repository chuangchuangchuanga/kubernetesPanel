import request from "@/utils/request.js";


export function getNamespaceList() {
    return request({
        url: "/namespace",
        method: "GET",
    })
}


