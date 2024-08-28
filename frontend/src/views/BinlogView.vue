<script setup lang="ts">
import { ref } from "vue";
import LogsList from "./binlog/LogsList.vue";
import {
  GetSystemFile,
  GetDecodeRowCount,
  SaveToSqlite, ClearAllData,
} from "@wailsjs/go/controllers/Files";
import { ElMessageBox } from "element-plus";
import { fa, tr } from "element-plus/es/locale/index.mjs";

const logCount = ref(0);
const dialogVisible = ref(false);
const percentage = ref(0);
const colors = [
  { color: "#f56c6c", percentage: 20 },
  { color: "#e6a23c", percentage: 40 },
  { color: "#5cb87a", percentage: 60 },
  { color: "#1989fa", percentage: 80 },
  { color: "#6f7ad3", percentage: 100 },
];

const handleClose = (done: () => void) => {
  ElMessageBox.confirm("退出弹窗, 解析不会中断！").then(() => {
    done();
  });
};

const childRef = ref();

const submitClear = () => {
  ElMessageBox.confirm("清空所有数据?").then(() => {
    ClearAllData().then(()=>{
      childRef.value.doSomething();
    })
  });
};

const submitUpload = () => {
  GetSystemFile().then((res) => {
    if (res != "") {
      percentage.value = 0;
      dialogVisible.value = true;

      var inteID = setInterval(() => {
        if (percentage.value <= 98) {
          percentage.value += 1;
        } else if (percentage.value >= 100) {
          dialogVisible.value = false;

          clearInterval(inteID);
        }
        GetDecodeRowCount().then((res) => {
          logCount.value = res;
        });
      }, 1000);

      SaveToSqlite(res).then((res) => {
        percentage.value = 100;
        dialogVisible.value = false;
        childRef.value.doSomething();
      });
    }
  });
};
</script>

<template>
  <div class="app-container">
    <div>
      <el-button type="primary" @click="submitUpload">
        Import Binlog File
      </el-button>

      <el-button type="primary" @click="submitClear">
        Clear ALL
      </el-button>
    </div>

    <LogsList ref="childRef" />

    <el-dialog
      v-model="dialogVisible"
      title="RUNNING。。。"
      :before-close="handleClose"
      center
    >
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
