<script lang="ts" setup>
import { inject, ref } from "vue";
import { GetTitleList } from "@wailsjs/go/controllers/Binlog";
import { db } from "@/store/db";
import { fa } from "element-plus/es/locale/index.mjs";

const FileList: any = ref([]);

function GetTitle() {
  if (db.ID == 0) {
    return;
  }

  GetTitleList(db.ID).then((res) => {
    FileList.value = res;
  });
}

GetTitle();

var selectedButton: any = ref(db.ID);

function clickButton(o: any) {
  selectedButton.value = o.ID;
  db.ID = o.ID;
}
</script>

<template>
  <div>
    <div v-for="o in FileList" :key="o">
      <el-button
        class="item"
        @click="clickButton(o)"
        :class="{ 'is-selected': selectedButton === o.ID }"
      >
        {{ o.Table }}
      </el-button>
    </div>
  </div>
</template>

<style scoped>
.item {
  box-shadow: 0 1px 3px hsla(0, 1%, 38%, 0.1);
  width: 100%;
  margin-bottom: 10px;
  margin-left: 0px;
  margin-right: 10px;
  padding: 5px;
  text-align: center;
  box-sizing: border-box;
}

.is-selected {
  background-color: #66c23a; /* 选中状态的背景色 */
  color: white; /* 选中状态的文字颜色 */
}
</style>
å
