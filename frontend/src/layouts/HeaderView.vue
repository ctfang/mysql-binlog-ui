<script lang="ts" setup></script>

<template>
  <div>
    <div class="header-logo">
      <div class="right-angle-triangle"></div>
    </div>

    <div>
      <span class="header-title">Mysql Binlog UI</span>
    </div>
  </div>
</template>

<style scoped>
.header-logo {
  position: absolute;
  right: 20px;
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

.right-angle-triangle {
  width: 20px;
  height: 20px;
  background: linear-gradient(
    45deg,
    rgb(231, 210, 210),
    orange,
    yellow,
    rgb(255, 255, 255),
    rgb(214, 214, 242),
    rgb(211, 195, 224),
    violet
  );
  transform: rotate(45deg);
}
</style>
