<template>
  <div class="overview">
    <div :style="{ display: 'flex' }">
      <el-card style="max-width: 360px">
        <template #header>
          <div class="card-header">
            <span>系统信息</span>
          </div>
        </template>
        <p>{{ data.info }}</p>
      </el-card>
      <el-card style="max-width: 360px">
        <template #header>
          <div class="card-header">
            <span>ip 信息</span>
          </div>
        </template>
        <p>{{ data.ip }}</p>
      </el-card>
      <el-card style="max-width: 360px">
        <template #header>
          <div class="card-header">
            <span>地域</span>
          </div>
        </template>
        <p>
        {{ data.location }}
        </p>
      </el-card>
      <el-card style="max-width: 360px">
        <template #header>
          <div class="card-header">
            <span>当前时间</span>
          </div>
        </template>
        <p>
        {{ currentTimestamp }}
        <br/>
        {{ currentTimeStr }}
        </p>
      </el-card>
    </div>
  </div>
</template>
<style scoped>
</style>
<script setup lang="ts">
import {onMounted, onUnmounted, reactive, ref} from 'vue'
import {IpInfo, LocationInfo, SystemInfo} from '../../../wailsjs/go/backend/App'

const currentTimestamp = ref(0); // 创建一个响应式变量来存储当前时间戳
const currentTimeStr = ref(''); // 创建一个响应式变量来存储当前时间字符串

// 函数用来更新时间戳
const updateTimestamp = () => {
  currentTimestamp.value = Date.now(); // 获取当前时间戳（毫秒）
  currentTimeStr.value = new Date(Date.now()).toLocaleString('zh-CN', { hour12: false }); // 获取当前时间字符串
};

let interval: any;

const data = reactive({
  info: "",
  ip: "",
  location: "",
})

onMounted(() => {
    systemInfo()
    ipInfo()
    locationInfo()
    updateTimestamp();
    interval = setInterval(updateTimestamp, 1000); // 每秒更新一次时间戳
  }
)

onUnmounted(() => {
  clearInterval(interval); // 组件销毁时清除定时器
})

function systemInfo() {
  SystemInfo().then(result => {
    data.info = result
  })
}

function ipInfo() {
  IpInfo().then(result => {
    data.ip = result
  })
}

function locationInfo() {
  LocationInfo().then(result => {
    data.location = result
  })
}
</script>
