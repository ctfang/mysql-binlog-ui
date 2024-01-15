<script lang="ts" setup>
import { RouterView } from "vue-router";
import IndexView from "@/layouts/index.vue";
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
  <div id="app">
    <IndexView>
      <RouterView />
    </IndexView>
  </div>
</template>
