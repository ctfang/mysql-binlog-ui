<script lang="ts" setup>

import { Exit } from "@wailsjs/go/controllers/Help";
import {WindowFullscreen, WindowMinimise} from "@wailsjs/runtime";

</script>

<template>
  <div>
    <div class="header-logo">
      <div class="window-controls">
        <div class="button red" @click="Exit"></div>
        <div class="button yellow" @click="WindowMinimise"></div>
        <div class="button green" @click="WindowFullscreen"></div>
      </div>
    </div>

    <div>
      <span class="header-title">Mysql Binlog UI</span>
    </div>
  </div>
</template>

<style scoped>
.header-logo {
  position: absolute;
  left: 13px;
  height: 100%;
  box-sizing: border-box;
  padding-top: 9px;
}

.header-title {
  position: absolute;
  left: calc(50% - 50px);
  height: 100%;
  box-sizing: border-box;
  padding-top: 10px;
}

.window-controls {
  display: flex;
  gap: 8px;
  border-radius: 6px;
  width: fit-content;
  margin: 4px auto;
}

.window-controls .button {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background-color: red;
  border: 1px solid rgba(0, 0, 0, 0.1);
  cursor: pointer;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.window-controls .button.yellow {
  background-color: #fdbc40;
}

.window-controls .button.green {
  background-color: #34c84a;
}
</style>
