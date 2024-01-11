<script lang="ts" setup>
import { AppMenu } from '@/router/index'
</script>

<template>
  <div class="layout-sidebar" style="--wails-draggable: drag">
    <div v-for="item in AppMenu" :key="item.path">
      <div class="menu-item" v-if="item.meta">
        <router-link :to="item.path">
          <div class="menu-icon">
            <component :is="item.meta.icon"></component>
          </div>
        </router-link>
      </div>
    </div>

    <div class="menu-item bottom-div">
      <router-link to="/about">
        <div class="menu-icon">
          <setting />
        </div>
      </router-link>
    </div>
  </div>
</template>

<style scoped>
.menu-item {
  margin: 15px 10px 20px 12px;

  place-items: center;
  place-content: center;
  width: 38px;
  height: 38px;
  color: #bcacac;
  border: 1px solid rgba(60, 60, 60, 0.12);
  background: #ffffff;
  border-radius: 5px;

  box-shadow:
    -2.8px -9.1px 32.1px rgba(0, 0, 0, 0.02),
    -4.4px -14.4px 53px rgba(0, 0, 0, 0.028),
    -4.2px -13.8px 63.4px rgba(0, 0, 0, 0.035),
    -0.9px -2.9px 64.7px rgba(0, 0, 0, 0.042),
    8.8px 28.9px 61.2px rgba(0, 0, 0, 0.05),
    32px 105px 80px rgba(0, 0, 0, 0.07);
}

.menu-icon {
  box-sizing: border-box;
  padding-left: 7px;
  padding-top: 6px;
  width: 30px;
}
.bottom-div {
  position: absolute;
  bottom: 0px;
}
</style>
