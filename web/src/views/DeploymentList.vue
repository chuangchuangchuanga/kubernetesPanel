<template>
  <el-table :data="deploymentList" stripe style="width: 100%">
    <el-table-column prop="Name" label="Name" width="180" />
    <el-table-column prop="Namespace" label="namespace" />
  </el-table>
</template>

<script >
import {getDeploymentList} from "@/api/api.js";
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
    }
  },
  mounted() {
    this.loadData()
  }
}
</script>