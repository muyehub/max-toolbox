<template>
  <div class="app-container">
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    @select="handleSelect"
  >
    <el-menu-item v-for="item in menuItems" :index="item.key" :key="item.key">
      <RouterLink :to="item.key">{{ item.label }}</RouterLink>
    </el-menu-item>
  </el-menu>
  <RouterView></RouterView>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted} from 'vue'
import {RouterView, useRoute, useRouter, RouterLink} from 'vue-router'

const router = useRouter()
const route = useRoute()

const activeIndex = computed(() => {
  return route.path
})

const menuItems = [
  { key: '/overview', label: '概览' },
  { key: '/tool', label: '工具' },
  { key: '/news', label: '新闻' }
]

const handleSelect = (key: string) => {
    router.push(key)
}

</script>

<style scoped>
.app-container {
  margin-top: 25px;
}
a,
.router-link,
.router-link-active {
  text-decoration: none;
}
</style>

