<template>
  <n-card style="height: 100%" content-style="height: 100%;padding: 10px">
    <n-scrollbar>
      <p v-for="log in logs" :key="log">
        {{ log.type }} {{ log.payload }}
      </p>
    </n-scrollbar>
  </n-card>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import {Events} from "@wailsio/runtime";

const logs = ref<any>([])

onMounted(() => {
  Events.On("logs", (data: any) => {
    logs.value.push(JSON.parse(data.data))
    // 仅保留最新100条
    if (logs.value.length > 100) {
      logs.value.shift()
    }
  })
})

</script>
<style scoped>

</style>