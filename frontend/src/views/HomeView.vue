<script lang="ts" setup>
import DbList from "./db/DbList.vue";
import DbDetails from "./db/DbDetails.vue";
import { useRouter } from "vue-router";
import { db } from "@/store/db";

const router = useRouter();

function goToDetails() {
  if (!db.ID) {
    router.push({ name: "binlogs" });
  }
}

goToDetails();
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
  width: 200px;
  padding: 10px;
  float: left;
}

.foter-container {
  width: calc(100% - 200px);
  box-shadow: 0 1px 3px hsla(0, 0%, 7%, 0.1);
  height: var(--main-height);
  padding: 0 0 10px 10px;
  background-color: #ffffff;
  box-sizing: border-box;
  float: right;
}
</style>
