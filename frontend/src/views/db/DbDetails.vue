<script lang="ts" setup>
import { reactive, inject } from 'vue'

const formInline = reactive({
  user: '',
  region: '',
  date: ''
})

const onSubmit = () => {
  console.log('submit!')
}

const tableData = [
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

function loadData() {
  let db = inject('db')

  console.log(db)
}
</script>

<template>
  <div>
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item>
        <el-input v-model="formInline.user" placeholder="table name" clearable />
      </el-form-item>
      <el-form-item>
        <el-select v-model="formInline.region" placeholder="Activity zone" clearable>
          <el-option label="Zone one" value="shanghai" />
          <el-option label="Zone two" value="beijing" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-date-picker v-model="formInline.date" type="date" placeholder="Pick a date" clearable />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Query</el-button>
      </el-form-item>
    </el-form>

    <br />

    <el-table :data="tableData" style="width: 100%">
      <el-table-column prop="date" label="导入时间" width="180" />
      <el-table-column prop="name" label="文件名称" width="180" />
      <el-table-column prop="address" label="文件地址" />
    </el-table>
  </div>
</template>
