<template>
  <el-container class="layout-container-demo" >
    <el-aside width="200px">
      <el-scrollbar>
        <div class="logo">
          <router-link :to="{name: 'Kubernetes'}">
            <img class="logoStyle" src="../assets/logo.png"></img>
          </router-link>
        </div>
        <el-menu>
          <router-link :to="{ name: 'DeploymentLists' }">
              <el-menu-item index="2" >Deployment</el-menu-item>
          </router-link>
        </el-menu>
      </el-scrollbar>
    </el-aside>

    <el-container>
      <el-header style="text-align: right; font-size: 12px">
        <div class="toolbar">
          <div>NameSpace:</div>
          <div class="flex flex-wrap gap-4 items-center">
            <el-select v-model="deploymentSelect" placeholder="namespace"  size="large" @change="namespaceSelectEvent"  style="width: 240px">
              <el-option v-for="item in deploymentListData" :key="item" :label="item" :value="item"/>
            </el-select>
          </div>
        </div>
      </el-header>

      <el-main>
        <el-scrollbar>

          <router-view />

        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>



<script>
import { getNamespaceList } from '@/api/api';
import { namespaceSelectStore } from "@/stores/namespaceSelect.js";


const useNamespaceSelect1 = namespaceSelectStore();

export default {
  name: 'kubernetesPanel',
  data() {
    return {
      deploymentSelect : '',
      deploymentListData : [],
    }
  },

  methods: {
    namespaceSelectEvent(value){
      useNamespaceSelect1.changeSelectedNamespace(value);
      window.location.reload();
    },

    async getNamespaceList(){
      try {
        const response = await getNamespaceList();
        this.deploymentListData = response.data.data;
      }catch (error) {
        console.log(error);
      }
    }
  },
  mounted() {
    this.deploymentSelect = useNamespaceSelect1.selectedNamespace;
    console.log("mounted拿到的namespace",useNamespaceSelect1.selectedNamespace);
    // 赋值给 data 中的 deploymentSelect
    this.getNamespaceList();

  }
}
</script>


<style scoped>
.layout-container-demo .el-header {
  position: relative;
  background-color: var(--el-color-primary-light-7);
  color: var(--el-text-color-primary);
}
.layout-container-demo .el-aside {
  color: var(--el-text-color-primary);
  background: var(--el-color-primary-light-8);
}
.layout-container-demo .el-menu {
  border-right: none;
}
.layout-container-demo .el-main {
  padding: 0;
}
.layout-container-demo .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}

.logo{
  height: 60px;
}
a {
  text-decoration: none;
}
</style>