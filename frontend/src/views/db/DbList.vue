<script lang="ts" setup>
import { inject, type Ref } from 'vue'

let db = inject<Ref<any>>('db')!

const FileList = [
  {
    name: '23-01',
    size: '100kb',
    type: 'image/jpeg',
    url: '/static/1.jpg'
  },
  {
    name: '24-05',
    size: '100kb',
    type: 'image/jpeg',
    url: '/static/1.jpg'
  }
]

function clickButton(o: any) {
  db.value = o
}
</script>

<template>
  <div>
    <el-button v-for="o in FileList" :key="o" class="item" @click="clickButton(o)">
      {{ o.name }}
    </el-button>
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
}
</style>
å
