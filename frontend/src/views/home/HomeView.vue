<template>
  <div class="container">
    <n-card content-style="padding: 10px" style="margin-bottom: 5px">
      <n-flex justify="space-between">
        <n-radio-group v-model:value="appStore.proxyMode" @update-value="changeProxyMode">
          <n-radio-button
              v-for="song in songs"
              :key="song.value"
              :value="song.value"
              :label="song.label"
          />
        </n-radio-group>
        <n-flex>
          <n-button type="primary" ghost @click="downloadTheKernel">下载内核</n-button>
          <n-button type="primary" @click="startSingBox" :disabled="appStore.isRunning">启动</n-button>
          <n-button type="error" @click="stopSingBox" :disabled="!appStore.isRunning">停止</n-button>
        </n-flex>
      </n-flex>
    </n-card>
    <n-card style="height: calc(100% - 60px)">
      <ChartView/>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import {AppService} from '@api/changeme/app/service'
import {ref} from "vue";
import {useMessage} from 'naive-ui'
import ChartView from "@/views/home/components/ChartView.vue";
import useAppStore from "@/stores/app/AppStore.ts";
import {useEventsStore} from "@/stores/events/EventsStore.ts";

const message = useMessage()
const appStore = useAppStore()
const events = useEventsStore()
const songs = ref([
  {
    value: 'system',
    label: '系统代理'
  },
  {
    value: 'tun',
    label: 'Tun'
  }
])

// 切换代理模式
function changeProxyMode(mode: string) {
  AppService.ChangeProxyMode(mode).then(() => {
    message.success("切换成功")
  })
}

function downloadTheKernel() {
  AppService.DownloadLatestKernel().then(() => {
    message.success("下载完成")
  })
}

function startSingBox() {
  AppService.StartCommand().then(() => {
    appStore.isRunning = true
    message.success("运行成功")
    // 延迟5s
    setTimeout(() => {
      events.listen()
    }, 5000)
  })
}

function stopSingBox() {
  AppService.StopCommand().then(() => {
    appStore.isRunning = false
    message.success("停止成功")
  })
}
</script>

<style scoped>

</style>
