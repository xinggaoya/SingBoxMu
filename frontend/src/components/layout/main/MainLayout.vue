<template>
  <div style="height: 100%;padding: 10px">
    <div style="height: 80px">
      <HeaderLayout/>
    </div>
    <div style="height: calc(100% - 100px);display: flex">
      <div style="height: 100%;width: 200px;margin-right: 5px">
        <LayoutMenu/>
      </div>
      <div style="height: 100%;flex: 1">
        <RouterView/>
      </div>
    </div>
  </div>

</template>

<script lang="ts" setup>

import LayoutMenu from "@/components/layout/menu/LayoutMenu.vue";
import HeaderLayout from "@/components/layout/header/HeaderLayout.vue";
</script>

<style scoped>
</style>
