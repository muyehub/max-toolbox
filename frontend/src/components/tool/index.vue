<template>
  <el-container>
    <el-aside>
      <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
        <el-radio-button :value="false">expand</el-radio-button>
        <el-radio-button :value="true">collapse</el-radio-button>
      </el-radio-group>
      <el-menu
          class="el-menu-vertical-demo"
          :collapse="isCollapse"
          @open="handleOpen"
          @close="handleClose"
          :router="true"
          :default-active="$route.path"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon>
              <location/>
            </el-icon>
            <span>base64</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="/tool/base64">base64转换</el-menu-item>
          </el-menu-item-group>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-main>
      <RouterView></RouterView>
    </el-main>
  </el-container>
</template>
<script lang="ts" setup>
import {ref} from 'vue';
import {RouterView} from 'vue-router'

import {Location,} from '@element-plus/icons-vue'

const isCollapse = ref(false)
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>
<style scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
