<script setup lang="ts">
import { OpenURL } from "@wailsjs/go/controllers/Help";

function goto(url: string) {
  OpenURL(url);
}
</script>

<template>
  <div class="about app-container">
    <h1>MySQL Binlog UI</h1>
    <p>
      读取原始的 MySQL Binlog文件后, 生成更加可读的结构化文件。<br />
      解析内容后，可以在 UI 上直接便捷的查看。 也可以通过生成的 sqlite
      数据库文件，直接在数据库客户端上查看。<br />
    </p>
    <p>
      Created by
      <a href="" @click="goto('https://github.com/ctfang')"> ctfang </a>
    </p>

    <p>
      Star it on GitHub:
      <el-link @click="goto('https://github.com/ctfang/mysql-binlog-ui')">
        https://github.com/ctfang/mysql-binlog-ui
      </el-link>
    </p>
  </div>
</template>

<style scoped></style>
