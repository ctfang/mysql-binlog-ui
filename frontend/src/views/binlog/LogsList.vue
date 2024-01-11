<script lang="ts" setup>
import { ref } from 'vue'
let tableData: any[] = []

function loadData() {
  tableData = [
    {
      date: '2016-05-03',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles'
    },
    {
      date: '2016-05-02',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles'
    },
    {
      date: '2016-05-04',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles'
    },
    {
      date: '2016-05-01',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles'
    }
  ]
}
</script>

<template>
  <el-table :data="tableData" style="width: 100%">
    <el-table-column prop="date" label="导入时间" width="180" />
    <el-table-column prop="name" label="文件名称" width="180" />
    <el-table-column prop="address" label="文件地址" />
  </el-table>
</template>
