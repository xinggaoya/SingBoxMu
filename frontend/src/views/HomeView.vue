<template>
  <div class="container">
      <n-flex>
        <n-radio-group @update-value="changeProxyMode">
          <n-radio-button
              v-for="song in songs"
              :key="song.value"
              :value="song.value"
              :label="song.label"
          />
        </n-radio-group>
        <n-button type="primary" ghost @click="downloadTheKernel">下载内核</n-button>
        <n-button type="primary" @click="startSingBox">启动</n-button>
        <n-button type="error" @click="stopSingBox">停止</n-button>
      </n-flex>
  </div>
</template>

<script setup lang="ts">
import {AppService} from '@api/changeme/app/service/index'
import {onMounted, ref} from "vue";
import {useMessage} from 'naive-ui'
import {Events} from "@wailsio/runtime";

const message = useMessage()
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

onMounted(()=>{
  Events.On("stdout", (data:string) => {
    console.log(data)
  })
  Events.On("stderr", (data:string) => {
    console.log(data)
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
