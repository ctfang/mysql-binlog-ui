<script lang="ts" setup>
import DbList from './db/DbList.vue'
import DbDetails from './db/DbDetails.vue'
import { provide, ref } from 'vue'

let db = ref({})

provide('db', db)
</script>

<template>
  <div class="db-main">
    <div class="db-title">
      <DbList />
    </div>
    <div class="foter-container">
      <DbDetails />
    </div>
  </div>
</template>

<style scoped>
.db-main {
  height: var(--main-height);
  display: flex;
}

.db-title {
  height: var(--main-height);
  width: 200px;
  padding: 10px;
  float: left;
}

.foter-container {
  width: calc(100% - 200px);
  box-shadow: 0 1px 3px hsla(0, 0%, 7%, 0.1);
  height: var(--main-height);
  padding: 10px;
  background-color: #ffffff;
  box-sizing: border-box;
  float: right;
}
</style>
