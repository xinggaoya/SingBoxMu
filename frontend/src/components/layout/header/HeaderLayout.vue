<script setup lang="ts">
import {BulbOutline, ChevronDown, CloseCircleOutline, Code,} from '@vicons/ionicons5'
import logo from "@/assets/app-icon.png";
import useAppStore from "@/stores/app/AppStore.ts";
import {Window} from "@wailsio/runtime";

const appStore = useAppStore()
const AppName = import.meta.env.VITE_APP_TITLE

function handleDark() {
  appStore.toggleDark()
}

async function handelCloseWindow() {
  await Window.Hide()
}

async function handelHideWindow() {
  await Window.Hide()
}

async function handelFullScreen() {
  await Window.ToggleFullscreen()
}


</script>

<template>
  <div class="layout-header" style="--wails-draggable:drag">
    <n-space justify="space-between">
      <n-flex>
        <n-image :src="logo" width="32" height="32"/>
        <n-text class="header-text" strong>{{ AppName }}</n-text>
      </n-flex>
      <n-flex>

        <n-tooltip placement="bottom" trigger="hover">
          <template #trigger>
            <n-button @click="handleDark" strong secondary circle type="default">
              <template #icon>
                <n-icon>
                  <BulbOutline/>
                </n-icon>
              </template>
            </n-button>
          </template>
          <span>调整主题</span>
        </n-tooltip>
        <n-tooltip placement="bottom" trigger="hover">
          <template #trigger>
            <n-button @click="handelFullScreen" strong secondary circle type="default" class="pc-but">
              <template #icon>
                <n-icon>
                  <Code/>
                </n-icon>
              </template>
            </n-button>
          </template>
          <span>全屏</span>
        </n-tooltip>
        <n-tooltip placement="bottom" trigger="hover">
          <template #trigger>
            <n-button @click="handelHideWindow" strong secondary circle type="default" class="pc-but">
              <template #icon>
                <n-icon>
                  <ChevronDown/>
                </n-icon>
              </template>
            </n-button>
          </template>
          <span>最小化</span>
        </n-tooltip>
        <n-tooltip placement="bottom" trigger="hover">
          <template #trigger>
            <n-button @click="handelCloseWindow" strong secondary circle type="default" class="pc-but">
              <template #icon>
                <n-icon>
                  <CloseCircleOutline/>
                </n-icon>
              </template>
            </n-button>
          </template>
          <span>关闭窗口</span>
        </n-tooltip>
      </n-flex>
    </n-space>
  </div>
</template>

<style scoped>

.header-text {
  font-weight: bold;
  line-height: 32px;
  font-size: 20px;
}

@media screen and (max-width: 768px) {
  .pc-but {
    display: none;
  }
}
</style>
