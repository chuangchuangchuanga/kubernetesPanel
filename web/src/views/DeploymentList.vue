<template>
  <el-table :data="deploymentList" stripe style="width: 100%">
    <el-table-column prop="Name" label="Name" width="380" />
    <el-table-column prop="Namespace" label="namespace" />
    <el-table-column label="操作" >
      <template #default="scope">
        <el-button link type="primary" size="small" @click="deploymentRestart(scope.row)">
          滚动重启
        </el-button>
        <el-button link type="primary" size="small" @click="clickDeploymentPodList(scope.row)">Detail</el-button>
      </template>
    </el-table-column>
  </el-table>


  <el-dialog v-model="podListdialogVisible"
      title="Pod列表"
      width="800"
  >
    <el-table :data="deploymentPodList" style="width: 100%">
      <el-table-column prop="Name" label="PodName" width="280" />
      <el-table-column prop="address" label="动作" />
    </el-table>

    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="podListdialogVisible = false">
          Confirm
        </el-button>
      </div>
    </template>
  </el-dialog>


</template>

<script >
import {getDeploymentList, getDeploymentPodList, postRestartDeployment} from "@/api/api.js";
import {namespaceSelectStore} from "@/stores/namespaceSelect.js";

export default {
  name: 'DeploymentLists',
  data() {
    return {
      podListdialogVisible: false,
      deploymentList: [],
      deploymentPodList: [],
    }
  },
  methods: {
    async clickDeploymentPodList(row) {
      this.podListdialogVisible = true;
      try {
        const requestData = {
          namespaceName : namespaceSelectStore().selectedNamespace,
          deploymentName: row.Name,
        }
        const response = await getDeploymentPodList(requestData);
        this.deploymentPodList = response.data.data;
      }catch (error) {
        console.log(error);
      }
    },

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