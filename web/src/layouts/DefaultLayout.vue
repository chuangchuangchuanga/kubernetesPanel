<template>
  <el-container class="layout-container-demo" >
    <el-aside width="200px">
      <el-scrollbar>
        <el-menu :default-openeds="['1', '3']">
              <el-menu-item index="2-2">Option 2</el-menu-item>
        </el-menu>
      </el-scrollbar>
    </el-aside>

    <el-container>
      <el-header style="text-align: right; font-size: 12px">
        <div class="toolbar">
          <div class="flex flex-wrap gap-4 items-center">
            <el-select v-model="deploymentSelect" placeholder="namespace" size="large" style="width: 240px">
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



<script >
import { getNamespaceList } from '@/api/api';



export default {
  name: 'DeploymentList',
  data() {
    return {
      deploymentSelect : "",
      deploymentListData : [],
    }
  },

  methods: {
    async getNamespaceList(){
      try {
        const response = await getNamespaceList();
        this.deploymentListData = response.data.data;
        console.log("1111", response.data.data);
      }catch (error) {
        console.log(error);
      }
    }
  },
  mounted() {
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
</style>