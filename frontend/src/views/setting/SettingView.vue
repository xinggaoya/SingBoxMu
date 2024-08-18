<template>
  <n-card  style="height: 100%" content-style="height: 100%;">
    <n-scrollbar>
      <n-divider>内核版本 {{version}}</n-divider>
      <n-flex>
        <n-button type="primary" size="small" ghost @click="downloadTheKernel">下载最新内核</n-button>
        <n-progress
            type="line"
            status="success"
            v-if="percentage > 0"
            :percentage
            indicator-placement="inside"
            processing
        />
      </n-flex>
      <n-divider>配置</n-divider>
      <n-form :model="configStore.form" ref="formRef" label-placement="left">
        <n-form-item label="开机自启" path="autoStart">
          <n-switch v-model:value="configStore.form.autoStart" @update-value="updateAutoStart"/>
        </n-form-item>
        <n-form-item label="内核自启" path="autoRun">
          <n-switch v-model:value="configStore.form.autoRun"/>
        </n-form-item>
      </n-form>
    </n-scrollbar>
  </n-card>
</template>

<script setup lang="ts">
import {useConfigStore} from "@/stores/config/ConfigStore";
import {AppService} from "@api/changeme/app/service";
import {onMounted, ref} from "vue";
import {useMessage} from "naive-ui";
import {Events} from "@wailsio/runtime";
import {GetVersion} from "@/api/clash/ClashApi";

const configStore = useConfigStore()
const formRef = ref()
const message = useMessage()
const version = ref<any>()
const percentage = ref<number>(0)

onMounted(() => {
  GetVersion().then((res: any) => {
    if (res) {
      version.value = res.version
    }
  })
})

function downloadTheKernel() {
  AppService.DownloadLatestKernel().then((res) => {
    if (res.code === 10000) {
      message.info("开始下载")
      Events.On(res.data, (data: any) => {
        const info:{
          totalSize:number,
          progress:number
        } = JSON.parse(data.data)

        // 计算百分比
        percentage.value = Math.round(info.progress / info.totalSize * 100)

        if (info.progress>= info.totalSize){
          Events.Off(res.data)
          message.success("下载完成")
        }
      })
    } else {
      message.error(res.msg)
    }
  })
}

const updateAutoStart = (val: boolean) => {
  if (val) {
    AppService.SetAutoStart().then(() => {
      message.success("设置成功")
    })
  } else {
    AppService.RemoveAutoStart().then(() => {
      message.success("移除成功")
    })
  }
}
</script>

<style scoped>

</style>