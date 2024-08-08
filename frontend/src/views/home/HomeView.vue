<template>
  <div class="container">
    <n-card content-style="padding: 10px" style="margin-bottom: 5px">
      <n-flex justify="space-between">
        <n-radio-group @update-value="changeProxyMode">
          <n-radio-button
              v-for="song in songs"
              :key="song.value"
              :value="song.value"
              :label="song.label"
          />
        </n-radio-group>
        <n-flex>
          <n-button type="primary" ghost @click="downloadTheKernel">下载内核</n-button>
          <n-button type="primary" @click="startSingBox">启动</n-button>
          <n-button type="error" @click="stopSingBox">停止</n-button>
        </n-flex>
      </n-flex>
    </n-card>
    <div>
      <EchartsView/>
    </div>
    <n-card style="height: calc(100% - 270px)" content-style="height: 100%;padding: 10px">
      <n-scrollbar>
        <p v-for="log in logs" :key="log">
          {{ log.type }} {{ log.payload }}
        </p>
      </n-scrollbar>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import {AppService} from '@api/changeme/app/service'
import {onMounted, ref} from "vue";
import {useMessage} from 'naive-ui'
import {Events} from "@wailsio/runtime";
import EchartsView from "@/views/home/components/EchartsView.vue";

const message = useMessage()
const logs = ref<any>([])
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

onMounted(() => {
  Events.On("logs", (data: any) => {
    logs.value.push(JSON.parse(data.data))
    // 仅保留最新100条
    if (logs.value.length > 100) {
      logs.value.shift()
    }
  })
})

// 切换代理模式
function changeProxyMode(mode: string) {
  AppService.ChangeProxyMode(mode).then(() => {
    message.success("操作成功")
  })
}

function downloadTheKernel() {
  AppService.DownloadLatestKernel().then(() => {
    message.success("操作成功")
  })
}

function startSingBox() {
  AppService.StartCommand().then(() => {
    message.success("操作成功")
  })
}

function stopSingBox() {
  AppService.StopCommand().then(() => {
    message.success("操作成功")
  })
}
</script>

<style scoped>

</style>
