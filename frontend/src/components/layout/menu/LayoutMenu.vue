<template>
  <n-card content-style="padding: 0;">
    <n-menu
        :options="menuOptions"
        :render-icon="renderIcon"
        :render-label="renderLabel"
        v-model:value="currentMenu"
        :on-update:value="onSelect"
    >
    </n-menu>
  </n-card>
</template>

<script setup lang="ts">
import {h, ref} from 'vue'
import {useRouter} from 'vue-router'
import {NEllipsis, NIcon} from 'naive-ui'
import {ChatbubbleOutline} from '@vicons/ionicons5'

const router = useRouter()
const currentMenu = ref(0)
const menuOptions = [
  {
    label: '概括',
    key: 0,
  },
  {
    label: '订阅',
    key: 1,
  },
  {
    label: '设置',
    key: 2,
    disabled: true,
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
      router.push('/setting')
      break
  }
  currentMenu.value = key
}

function renderIcon() {
  return h(NIcon, null, {default: () => h(ChatbubbleOutline)})
}

function renderLabel(option: any) {
  return h(NEllipsis, {
    class: 'menu-label',
  }, option.label)
}

</script>

<style lang="scss" scoped>

.n-menu {

  ::v-deep(.n-menu-item-content-header) {
    display: flex;
    justify-content: space-between;
  }

  ::v-deep(.menu-delete-button) {
    visibility: hidden;
  }

  ::v-deep( .n-menu-item:hover .menu-delete-button ) {
    visibility: initial;
  }

}

</style>
