<script setup lang="ts">
import { ref } from 'vue'
import { genFileId } from 'element-plus'
import LogsList from './binlog/LogsList.vue'
import { GetSystemFile, GetDecodeRowCount } from '@wailsjs/go/controllers/Files'
import { ElMessageBox } from 'element-plus'
import { fa, tr } from 'element-plus/es/locale/index.mjs'

const logCount = ref(0)
const dialogVisible = ref(false)
const percentage = ref(0)
const colors = [
  { color: '#f56c6c', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#5cb87a', percentage: 60 },
  { color: '#1989fa', percentage: 80 },
  { color: '#6f7ad3', percentage: 100 }
]

const handleClose = (done: () => void) => {
  ElMessageBox.confirm('退出弹窗, 解析不会中断！').then(() => {
    done()
  })
}

const submitUpload = () => {
  dialogVisible.value = true

  GetSystemFile().then((res) => {
    percentage.value = 100
    if (res != '') {
      dialogVisible.value = false
      ElMessageBox.alert(res, '执行失败', {
        type: 'warning',
        confirmButtonText: '确定'
      })
    }
  })

  var inteID = setInterval(() => {
    if (percentage.value <= 98) {
      percentage.value += 1
    } else if (percentage.value >= 100) {
      dialogVisible.value = false

      clearInterval(inteID)
    }
    GetDecodeRowCount().then((res) => {
      logCount.value = res
    })
  }, 1000)
}
</script>

<template>
  <div class="app-container">
    <div>
      <el-button type="primary" @click="submitUpload">select file</el-button>
    </div>

    <LogsList />

    <el-dialog v-model="dialogVisible" title="解析中。。。" :before-close="handleClose" center>
      <div class="demo-progress">
        <el-progress
          :text-inside="true"
          :percentage="percentage"
          :stroke-width="26"
          :color="colors"
        />
        <span>decode row = {{ logCount }}</span>
      </div>
    </el-dialog>
  </div>
</template>

<style scoped>
.demo-progress .el-progress--line {
  margin-bottom: 15px;
}
</style>
