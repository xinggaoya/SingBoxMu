<template>
  <div>
    <n-card content-style="padding: 0;">
      <n-menu
          :options="menuOptions"
          v-model:value="currentMenu"
          :on-update:value="onSelect"
      >
      </n-menu>
    </n-card>
    <n-card style="margin-top: 5px">
      <n-flex vertical>
        <n-tag type="primary">
          <span>使用内存：{{ events.memory.inuse }}Mb</span>
        </n-tag>
        <n-tag type="success">
          <span>内存限制：{{ events.memory.oslimit }}Mb</span>
        </n-tag>
        <n-tag type="warning">
          <span>总流量：{{ totalUseTraffic }}Mb</span>
        </n-tag>
      </n-flex>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import {computed, h, ref} from 'vue'
import {useRouter} from 'vue-router'
import {NIcon} from 'naive-ui'
import {AtCircleOutline, BarChartOutline, InformationCircleOutline, SettingsOutline} from '@vicons/ionicons5'
import {useEventsStore} from "@/stores/events/EventsStore.ts";

const router = useRouter()
const currentMenu = ref(0)
const events = useEventsStore()
// 总流量
const totalUseTraffic = computed(() => {
  return events.totalUseTraffic.toFixed(2)
})
const menuOptions = [
  {
    label: '概括',
    key: 0,
    icon: () => {
      return h(NIcon, null, {default: () => h(BarChartOutline)})
    },
  },
  {
    label: '订阅',
    key: 1,
    icon: () => {
      return h(NIcon, null, {default: () => h(AtCircleOutline)})
    }
  },
  {
    label: '日志',
    key: 2,
    icon: () => {
      return h(NIcon, null, {default: () => h(InformationCircleOutline)})
    }
  },
  {
    label: '设置',
    key: 3,
    icon: () => {
      return h(NIcon, null, {
        default: () => h(SettingsOutline)
      })
    }
  },
]


function onSelect(key: number) {
  switch (key) {
    case 0:
      router.push('/')
      break
    case 1:
      router.push('/sub')
      break
    case 2:
      router.push('/log')
      break
    case 3:
      router.push('/setting')
      break
    default:
      break
  }
  currentMenu.value = key
}

</script>

<style lang="scss" scoped>

</style>
