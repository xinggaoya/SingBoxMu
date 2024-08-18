<template>
  <n-card title="订阅" class="about">
    <n-flex>
      <n-input v-model:value="url" placeholder="请输入订阅地址"/>
      <n-button type="primary" ghost @click="downloadSub" :loading>下载订阅</n-button>
    </n-flex>
  </n-card>
</template>
<script setup lang="ts">
import {AppService} from '@api/changeme/app/service'
import {ref} from "vue";
import {useMessage} from 'naive-ui'
import useAppStore from "@/stores/app/AppStore";

const url = ref('')
const message = useMessage()
const loading = ref(false)
const appState = useAppStore()

// 下载订阅
function downloadSub() {
  if (url.value === '') {
    message.error("订阅地址不能为空")
    return
  }
  loading.value = true
  AppService.DownloadSubscription(url.value).then(() => {
    loading.value = false
    message.success("下载成功")
    appState.proxyMode = ''
  })
}
</script>
<style lang="scss" scoped>

.about {
  height: 100%;
  overflow: auto;
}
</style>
