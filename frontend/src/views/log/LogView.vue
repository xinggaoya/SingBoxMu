<template>
  <n-card style="height: 100%" content-style="height: 100%;padding: 10px">
    <n-scrollbar>
      <n-flex vertical>
        <div v-for="(log,index) in events.logs" :key="index">
          <n-tag :type="log.type"> {{ log.type }}</n-tag>
          {{ log.payload }}
        </div>
      </n-flex>
    </n-scrollbar>
  </n-card>
</template>

<script setup lang="ts">
import {onBeforeUnmount} from "vue";
import {Events} from "@wailsio/runtime";
import {useEventsStore} from "@/stores/events/EventsStore";

const events = useEventsStore()

onBeforeUnmount(() => {
  Events.Off("logs");  // 清除事件监听器
});


</script>
<style scoped>

</style>