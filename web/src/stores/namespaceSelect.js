import { defineStore,storeToRefs } from "pinia";



export const namespaceSelectStore = defineStore("namespaceSelectStore", {
    state: ()=> {
        return {
            selectedNamespace: "default"
        }
    },
    actions: {
       changeSelectedNamespace(namespace){
           console.log("传进来的值是:",namespace)
           this.selectedNamespace = namespace
       },
    },
    persist: {
        enabled: true,           // 启用持久化
        storage: localStorage,   // 可以使用 localStorage 或 sessionStorage
        paths: ['selectedNamespace'],     // 仅持久化 username 状态
    },
})