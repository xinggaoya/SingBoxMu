<template>
  <div>
    <n-menu
        :collapsed-width="64"
        :collapsed-icon-size="22"
        :options="menuStore.menuList"
        :render-icon="renderIcon"
        :render-label="renderLabel"
        :value="menuStore.currentMenu"
        :on-update:value="onSelect"
    >
    </n-menu>
    <n-space vertical v-if="showMenu">
      <n-skeleton height="40px" v-for="i in 10" :key="i" style="margin: 0 10px"/>
    </n-space>
    <n-empty description="你什么也找不到" v-if="!showMenu&&menuStore.menuList.length==0" style="margin: 20px">
      <template #extra>
        <n-button size="small" @click="createChat">
          新建一个新的话题吧
        </n-button>
      </template>
    </n-empty>
  </div>
</template>

<script setup lang="ts">
import {h, onMounted, ref} from 'vue'
import {useMenuStore} from '@/stores/menu/menuStore'
import {useRouter} from 'vue-router'
import {NButton, NEllipsis, NIcon, NSpace} from 'naive-ui'
import {ChatbubbleOutline} from '@vicons/ionicons5'

const menuStore = useMenuStore()
const router = useRouter()
const showMenu = ref(true)

onMounted(() => {
  menuStore.initMenu().then(() => {
    showMenu.value = false
  })
})

function createChat() {
  showMenu.value = true
}

function onSelect(key: number) {
  menuStore.currentMenu = key
  menuStore.setRoleNameByKey(key)
  router.push({path: `/chat/${key}`})
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
