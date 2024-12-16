<template>
  <el-table :data="deploymentList" stripe style="width: 100%">
    <el-table-column prop="Name" label="Name" width="380" />
    <el-table-column prop="Namespace" label="namespace" />
    <el-table-column label="操作" >
      <template #default="scope">
        <el-button link type="primary" size="small" @click="deploymentRestart(scope.row)">
          滚动重启
        </el-button>
        <el-button link type="primary" size="small">Detail</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script >
import {getDeploymentList, postRestartDeployment} from "@/api/api.js";
import {namespaceSelectStore} from "@/stores/namespaceSelect.js";

export default {
  name: 'DeploymentLists',
  data() {
    return {
      deploymentList: [],
    }
  },
  methods: {
    async loadData() {
      try {
        const  namespaceSelect = {
          namespaceName: namespaceSelectStore().selectedNamespace,
        }
        const  response = await getDeploymentList(namespaceSelect);
        this.deploymentList = response.data.data;
      }catch(err) {
        console.log(err)
      }
    },
    async deploymentRestart(row) {
      try {
        const requestData = {
          namespaceName : namespaceSelectStore().selectedNamespace,
          deploymentName: row.Name,
        }
        await postRestartDeployment(requestData);
        this.loadData()
      }catch (error) {
        console.log(error);
      }
    }
  },
  mounted() {
    this.loadData()
  }
}
</script>