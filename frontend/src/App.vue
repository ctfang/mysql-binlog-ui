<script lang="ts" setup>
import { RouterView } from 'vue-router'
import IndexView from '@/layouts/index.vue'
</script>

<template>
  <div id="app">
    <IndexView>
      <RouterView />
    </IndexView>
  </div>
</template>
