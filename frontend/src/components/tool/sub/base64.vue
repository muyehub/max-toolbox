<template>
  字符串:<el-input
      v-model="s"
      style="width: 240px"
      placeholder="Please input"
      clearable
  />
  <br/>
  <el-button type="primary" round @click="convertStr">转换为字符串</el-button>
  <br/>
  <el-button type="primary" round @click="convertBase64">转换为base64</el-button>
  <br/>
  base64:<el-input
      v-model="b"
      style="width: 240px"
      placeholder="Please input"
      clearable
  />
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {ConvertBase64ToStr,ConvertStrToBase64} from '../../../../wailsjs/go/backend/App'

let b = ref("")
let s = ref("")

function convertStr() {
  ConvertBase64ToStr(b.value).then(result => {
    s.value = result
  })
}

function convertBase64() {
  ConvertStrToBase64(s.value).then(result => {
    b.value = result
  })
}
</script>
<style scoped>

</style>