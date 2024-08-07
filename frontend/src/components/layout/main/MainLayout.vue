<template>
  <n-layout class="container" style="padding: 5px">
    <n-layout-header class="rounded-r">
      <n-card class="rounded-r" content-style="padding: 10px">
        <HeaderLayout/>
      </n-card>
    </n-layout-header>
    <n-layout has-sider class="container" content-style="height: 100%"
              style="padding: 5px 0;height: calc(100% - 60px)">
      <n-layout-sider
          class="container"
          collapse-mode="width"
          collapsed-width="0"
          :width="260"
          show-trigger="bar"
          style="margin-right: 5px"
          :native-scrollbar="false"
      >
        <n-card class="rounded-r" content-style="padding: 0">
          <template v-slot:default>
            <LayoutMenu/>
          </template>
        </n-card>
      </n-layout-sider>
      <n-layout content-style="height: 100%;"
                :native-scrollbar="false">
        <RouterView/>
      </n-layout>
    </n-layout>
  </n-layout>
</template>

<script lang="ts" setup>
import LayoutMenu from '@/components/layout/menu/LayoutMenu.vue'
import HeaderLayout from "@/components/layout/header/HeaderLayout.vue";

</script>

<style scoped>

.header-text {
  font-weight: bold;
  font-size: large;
  line-height: 32px;
}

@media screen and (max-width: 768px) {
}
</style>
