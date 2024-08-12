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

const appStore = useAppStore()

onMounted(() => {
  getAppRunStatus()
})

// 获取程序运行状态
function getAppRunStatus() {
  AppService.IsRunningAsAdmin().then((res) => {
    appStore.isAdminRun = res.code === 10000;
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
