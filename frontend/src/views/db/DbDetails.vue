<script lang="ts" setup>
import { reactive, watch, ref } from "vue";
import { db } from "@/store/db";
import { GetDetailsList } from "@wailsjs/go/controllers/Binlog";

const formInline: any = reactive({
  Table: "",
  Event: "",
  Text: "",
  Date: ref<[any, any]>,

  Page: 1,
  Limit: 13,
});

watch(db, function () {
  loadData();
});

const onSubmit = () => {
  loadData();
};

var tableData = ref([]);
var total = ref(0);

function loadData() {
  if (db.ID == 0) {
    return;
  }

  GetDetailsList(db.ID, formInline).then((res) => {
    tableData.value = res.List;
    total.value = res.Total;
  });
}

function dateShow(row: any) {
  return new Date(row.Timestamp).toLocaleString();
}
// 返回第一个[ 和 , 号前的数据
function firstKeyShow1(row: any) {
  return getFirstKeyShow(row.Row1);
}
function firstKeyShow2(row: any) {
  return getFirstKeyShow(row.Row2);
}

function getFirstKeyShow(row: string) {
  // 找到第一个 '[' 的位置
  const startIndex = row.indexOf("[");

  // 找到第一个 ',' 的位置
  const endIndex = row.indexOf(",");

  // 如果 '[' 和 ',' 都存在且位置合理
  if (startIndex !== -1 && endIndex !== -1 && endIndex > startIndex) {
    // 返回 '[' 和 ',' 之间的字符串
    return row.substring(startIndex + 1, endIndex);
  }

  // 如果找不到 '[' 或 ','，或者 ',' 在 '[' 之前出现，返回空字符串
  return "";
}

loadData();

const currentPage = ref(1);
const pageSize = ref(13);
const small = ref(false);
const background = ref(false);
const disabled = ref(false);
const handleSizeChange = (val: number) => {
  formInline.Limit = val;
  loadData();
};
const handleCurrentChange = (val: number) => {
  formInline.Page = val;
  loadData();
};
</script>

<template>
  <div class="table-main">
    <el-form
      :inline="true"
      :model="formInline.Table"
      class="search demo-form-inline"
    >
      <el-form-item>
        <el-input
          with="width: 100px"
          v-model="formInline.Table"
          placeholder="table name"
          clearable
        />
      </el-form-item>
      <el-form-item>
        <el-select v-model="formInline.Event" placeholder="event" clearable>
          <el-option label="insert" value="insert" />
          <el-option label="update" value="update" />
          <el-option label="delete" value="delete" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-input
          with="width: 100px"
          v-model="formInline.Text"
          placeholder="row text"
          clearable
        />
      </el-form-item>
      <el-form-item>
        <el-date-picker
          v-model="formInline.Date"
          type="datetimerange"
          range-separator="To"
          start-placeholder="Start date"
          end-placeholder="End date"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Query</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="tableData" style="width: 100%">
      <el-table-column type="expand">
        <template #default="props">
          <div m="4">
            <p m="t-0 b-2">Change Before: {{ props.row.Row1 }}</p>
            <p m="t-0 b-2">Change After_: {{ props.row.Row2 }}</p>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="ID" label="ID" width="80px" />
      <el-table-column
        prop="Table"
        label="Table"
        width="180"
        show-overflow-tooltip
      />
      <el-table-column prop="Event" label="Event" width="100" />
      <el-table-column
        prop="Timestamp"
        label="Timestamp"
        :formatter="dateShow"
        width="160"
      />
      <el-table-column
        prop="Row1"
        label="Before Forst Key"
        :formatter="firstKeyShow1"
        show-overflow-tooltip
      />
      <el-table-column
        prop="Row2"
        label="After Forst Key"
        :formatter="firstKeyShow2"
        show-overflow-tooltip
      />
    </el-table>
    <br />
    <div class="example-pagination-block">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[13, 50, 100, 200, 300, 400]"
        :small="small"
        :disabled="disabled"
        :background="background"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<style scoped>
.table-main {
  overflow: auto;
  height: var(--main-height);
}
/* 为滚动条轨道（整个滚动条背景）设置样式 */
.table-main::-webkit-scrollbar {
  width: 10px; /* 滚动条的宽度 */
  background-color: #f5f5f5; /* 滚动条轨道的背景色 */
}

/* 为滚动条滑块（可拖动部分）设置样式 */
.table-main::-webkit-scrollbar-thumb {
  background-color: #9e9e9e; /* 滚动条滑块的颜色 */
  border-radius: 5px; /* 滚动条滑块的圆角 */
  border: 2px solid #f5f5f5; /* 滚动条滑块的边框和轨道的背景色相同 */
}

/* 当鼠标悬停在滚动条滑块上时的样式 */
.table-main::-webkit-scrollbar-thumb:hover {
  background-color: #757575; /* 滚动条滑块的颜色（悬停时） */
}

.search {
  margin-top: 10px;
}
.example-pagination-block + .example-pagination-block {
  margin-top: 10px;
}
.example-pagination-block .example-demonstration {
  margin-bottom: 16px;
}
</style>
å
