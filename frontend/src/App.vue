<script setup lang="ts">
import {RouterView} from 'vue-router'
import {darkTheme, dateZhCN, zhCN} from "naive-ui";
import {computed} from "vue";
import useAppStore from "@/stores/app/AppStore.ts";

const appStore = useAppStore()

const theme = computed(() => {
  // 获取元素
  const element = document.getElementById('max-main-view');
  if (appStore.isDark) {
    element?.classList.remove('app-background');
    return darkTheme
  }
  element?.classList.add('app-background');
  return null
})

const themeOverrides = {
  "common": {
    "borderRadius": "5px",
    "primaryColor": "#409EFF",
    "primaryColorHover": "#66b1ff",
    "primaryColorPressed": "#3a8ee6",
    "primaryColorSuppl": "#409eff",
    "infoColor": "#909399",
    "infoColorHover": "#a6a9ad",
    "infoColorPressed": "#82848a",
    "infoColorSuppl": "#909399",
    "successColor": "#67c23a",
    "successColorHover": "#85ce61",
    "successColorPressed": "#5daf34",
    "successColorSuppl": "#67c23a",
    "warningColor": "#e6a23c",
    "warningColorHover": "#ebb563",
    "warningColorPressed": "#cf9236",
    "warningColorSuppl": "#e6a23c",
    "errorColor": "#f56c6c",
    "errorColorHover": "#f78989",
    "errorColorPressed": "#dd6161",
    "errorColorSuppl": "#f56c6c"
  }
}

</script>

<template>
  <n-config-provider :theme :theme-overrides="themeOverrides" :locale="zhCN" :date-locale="dateZhCN" class="container">
    <n-global-style/>
    <n-message-provider>
      <n-dialog-provider>
        <RouterView/>
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
</style>
