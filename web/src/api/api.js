import request from "@/utils/request.js";
import service from "@/utils/request.js";


export function getNamespaceList() {
    return service({
        url: "/namespace",
        method: "GET",
    })
}


