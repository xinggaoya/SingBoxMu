<template>
  <div class="about">
   <n-flex>
     <n-input v-model:value="url" placeholder="请输入订阅地址"/>
     <n-button type="primary" ghost @click="downloadSub">下载订阅</n-button>
   </n-flex>
  </div>
</template>
<script setup lang="ts">
import {AppService} from '@api/changeme/app/service/index'
import {ref} from "vue";
import {useMessage} from 'naive-ui'

const url = ref('')
const message = useMessage()

// 下载订阅
function downloadSub() {
  AppService.DownloadSubscription(url.value).then(() => {
    message.success("下载成功")
  })
}
</script>
<style lang="scss" scoped>

</style>
