<template>
  <div class="container">
    <n-card content-style="padding: 10px" style="margin-bottom: 5px">
      <n-flex justify="space-between">
        <n-radio-group v-model:value="appStore.proxyMode" @update-value="changeProxyMode">
          <n-radio-button
              value="system"
              label="系统代理"
          />
          <n-radio-button
              value="tun"
              label="Tun"
              :disabled="!appStore.isAdminRun"
          />
        </n-radio-group>
        <n-flex>
          <n-button type="primary" size="small" ghost @click="downloadTheKernel">下载内核</n-button>
          <n-button type="primary" size="small" v-if="!appStore.isAdminRun" ghost @click="restartAdminSingBox">
            以管理员重启
          </n-button>
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
import {onMounted} from "vue";
import {useMessage} from 'naive-ui'
import ChartView from "@/views/home/components/ChartView.vue";
import useAppStore from "@/stores/app/AppStore";
import {useEventsStore} from "@/stores/events/EventsStore";
import {useConfigStore} from "@/stores/config/ConfigStore";

const message = useMessage()
const appStore = useAppStore()
const config = useConfigStore()
const events = useEventsStore()

onMounted(() => {
  appStore.getKernelVersion().then(() => {
    if (config.form.autoRun && !appStore.isRunning) {
      if (!appStore.isAdminRun && appStore.proxyMode === 'tun') {
        changeProxyMode('system')
      }
      startSingBox()
    }
  })
})

// 切换代理模式
function changeProxyMode(mode: string) {
  if (mode === 'tun' && !appStore.isAdminRun) {
    appStore.proxyMode = 'system'
    message.error("请以管理员权限运行")
    return
  }
  AppService.ChangeProxyMode(mode).then((res) => {
    if (res.code === 10000) {
      message.success("切换成功")
    } else {
      message.error(res.msg)
    }
  })
}

function downloadTheKernel() {
  AppService.DownloadLatestKernel().then((res) => {
    if (res.code === 10000) {
      message.success("下载完成")
    } else {
      message.error(res.msg)
    }
  })
}

function restartAdminSingBox() {
  AppService.RestartAsAdmin().then((res) => {
    if (res.code === 10000) {
      message.success("重启成功")
    } else {
      message.error(res.msg)
    }
  })
}

function startSingBox() {
  if (!appStore.proxyMode) {
    message.error("请选择代理模式")
    return
  }
  AppService.StartCommand().then((res) => {
    if (res.code === 10000) {
      appStore.isRunning = true
      message.success("运行成功")
      events.listen()
    } else {
      message.error(res.msg)
    }
  })
}

function stopSingBox() {
  AppService.StopCommand().then((res) => {
    if (res.code === 10000) {
      appStore.isRunning = false
      events.memory.inuse = 0
      message.success("停止成功")
    } else {
      message.error(res.msg)
    }
  })
}
</script>

<style scoped>

</style>
