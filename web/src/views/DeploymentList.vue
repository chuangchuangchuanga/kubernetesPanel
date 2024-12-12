<template>
  <el-table :data="tableData" style="width: 100%">
    <el-table-column prop="date" label="序号" width="180" />
    <el-table-column prop="name" label="名称" width="180" />
    <el-table-column prop="address" label="状态" />
  </el-table>
</template>



<script >
import { getNamespaceList } from '@/api/api';
import {onMounted} from "vue";


const tableData = [
  {
    date: '2016-05-03',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
]



export default {
  name: 'DeploymentList',
  data() {
    return {
      deployment : [],
    }
  },

  methods: {
    async getNamespaceList(){
      try {
        getNamespaceList().then((responde) =>
        {
          this.deployment = responde.data
        },
        );
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