<template>
  <n-flex vertical class="height-100 max-main">
    <n-card content-style="padding: 10px;">
      <HeaderLayout/>
    </n-card>
    <n-flex style="height: calc(100% - 80px)">
      <n-card content-style="padding: 0;" style="width: 200px;">
        <LayoutMenu/>
      </n-card>
      <n-card content-style="padding: 10px;" style="flex: 1;">
        <RouterView/>
      </n-card>
    </n-flex>
  </n-flex>
</template>

<script lang="ts" setup>
import LayoutMenu from '@/components/layout/menu/LayoutMenu.vue'
import HeaderLayout from "@/components/layout/header/HeaderLayout.vue";
</script>

<style scoped>

.max-main {
  padding: 10px;
}

.height-100 {
  height: 100%;
}
</style>
