<template>
  <div class="max-card">
    <n-card
        id="max-main-view"
        class="app-background"
        content-style="height: 100%;padding: 10px;">
      <div style="height: 80px;margin-bottom: 5px">
        <HeaderLayout/>
      </div>
      <div style="height: calc(100% - 85px);display: flex">
        <div style="height: 100%;width: 200px;margin-right: 5px">
          <LayoutMenu/>
        </div>
        <div style="height: 100%;flex: 1">
          <RouterView/>
        </div>
      </div>
    </n-card>
  </div>

</template>

<script lang="ts" setup>
import LayoutMenu from "@/components/layout/menu/LayoutMenu.vue";
import HeaderLayout from "@/components/layout/header/HeaderLayout.vue";
import {useAppStore} from "@/stores/app/AppStore";
import {AppService} from "@api/changeme/app/service";
import {onMounted} from "vue";
import {useEventsStore} from "@/stores/events/EventsStore";
import {useConfigStore} from "@/stores/config/ConfigStore";
import {useMessage} from "naive-ui";

const appStore = useAppStore()
const message = useMessage()
const events = useEventsStore()
const config = useConfigStore()

onMounted(() => {
  getAppRunStatus()
})

// 获取程序运行状态
function getAppRunStatus() {
  AppService.IsRunningAsAdmin().then((res) => {
    appStore.isAdminRun = res.code === 10000;
  })
  AppService.IsRunning().then((res) => {
    appStore.isRunning = res.code === 10000;
    if (config.form.autoRun && !appStore.isRunning) {
      if (!appStore.isAdminRun && appStore.proxyMode === 'tun') {
        changeProxyMode('system')
      }
      startSingBox()
    }
  })
}

// 切换代理模式
function changeProxyMode(mode: string) {
  AppService.ChangeProxyMode(mode).then((res) => {
    if (res.code === 10000) {
      message.success("切换成功")
    } else {
      message.error(res.msg)
    }
  })
}


function startSingBox() {
  if (!appStore.proxyMode) {
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
</script>

<style scoped>
.max-card {
  height: 100%;
}

#max-main-view {
  height: 100%;
}
</style>
