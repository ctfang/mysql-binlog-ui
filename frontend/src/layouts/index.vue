<script lang="ts" setup>
import LeftMenu from './LeftMenu.vue'
import HeaderView from './HeaderView.vue'
</script>

<template>
  <div>
    <header class="layout-header" style="--wails-draggable: drag">
      <HeaderView />
    </header>
    <div class="layout-body">
      <LeftMenu />
      <main>
        <div class="layout-main">
          <div class="main">
            <slot></slot>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped></style>
