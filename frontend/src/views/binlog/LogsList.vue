<script lang="ts" setup>
import { GetTitleList, DeleteFile } from "@wailsjs/go/controllers/Files";
import { ref, defineExpose, type Ref, inject } from "vue";
import { ElMessageBox } from "element-plus";
import { db } from "@/store/db";
import { useRouter } from "vue-router";

const router = useRouter();

let tableData = ref([]);
let total = ref(0);

function loadData() {
  GetTitleList(1, 10).then((res) => {
    tableData.value = res.List;
    total.value = res.Total;
  });
}

loadData();

const handleEdit = (index: number, row: any) => {
  db.ID = row.ID;

  router.push("/");
};
const handleDelete = (index: number, row: any) => {
  ElMessageBox.confirm("Delete the file. Continue?", "Warning", {
    confirmButtonText: "OK",
    cancelButtonText: "Cancel",
    type: "warning",
  })
    .then(() => {
      DeleteFile(row.ID).then(() => {
        loadData();
      });
    })
    .catch(() => {});
};

function dateShow(row: any) {
  return new Date(row.Timestamp).toLocaleString();
}

function dateDatabase(row: any) {
  const lastIndex = row.Database.lastIndexOf("/");
  if (lastIndex === -1) {
    // 如果没有找到 '/'，返回整个路径
    return row.Database;
  }
  // 返回最后一个 '/' 之后的部分
  return row.Database.substring(lastIndex + 1);
}

const childBorder = ref(false);

const doSomething = () => {
  loadData();
};

defineExpose({ doSomething });
</script>

<template>
  <div>
    <el-table :data="tableData" style="width: 100%">
      <el-table-column type="expand">
        <template #default="props">
          <div m="4">
            <p m="t-0 b-2">File Form: {{ props.row.File }}</p>
            <p m="t-0 b-2">Database: {{ props.row.Database }}</p>
            <p m="t-0 b-2">Table: {{ props.row.Table }}</p>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="ID" label="ID" width="80px" />
      <el-table-column
        prop="Database"
        label="Database"
        :formatter="dateDatabase"
        min-width="100px"
      />
      <el-table-column prop="Table" label="Table" min-width="100px" />
      <el-table-column prop="FileSize" label="FileSize/MB" min-width="120px" />
      <el-table-column
        prop="Timestamp"
        label="Timestamp"
        :formatter="dateShow"
        width="160px"
      />

      <el-table-column label="Operations" width="170px">
        <template #default="scope">
          <el-button
            size="small"
            @click="handleEdit(scope.$index, scope.row)"
            type="success"
          >
            OpenDB
          </el-button>
          <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
          >
            Delete
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <br />
    <div class="example-pagination-block">
      <el-pagination layout="prev, pager, next" :total="total" />
    </div>
  </div>
</template>
