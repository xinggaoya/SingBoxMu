<template>
  <n-card  style="height: 100%" content-style="height: 100%;">
    <n-scrollbar>
      <n-divider>内核版本 {{version}}</n-divider>
      <n-form :model="configStore.form" ref="formRef" label-placement="left">
        <n-form-item label="开机自启" path="autoStart">
          <n-switch v-model:value="configStore.form.autoStart" @update-value="updateAutoStart"/>
        </n-form-item>
        <n-form-item label="内核自启" path="autoRun">
          <n-switch v-model:value="configStore.form.autoRun"/>
        </n-form-item>
      </n-form>
    </n-scrollbar>
  </n-card>
</template>

<script setup lang="ts">
import {useConfigStore} from "@/stores/config/ConfigStore";
import {AppService} from "@api/changeme/app/service";
import {onMounted, ref} from "vue";
import {useMessage} from "naive-ui";

const configStore = useConfigStore()
const formRef = ref()
const message = useMessage()
const version = ref<any>()

onMounted(() => {
  AppService.GetVersion().then((res: any) => {
    const date = JSON.parse(res.data)
    if (date) {
      version.value = date.version
    }
  })
})

const updateAutoStart = (val: boolean) => {
  if (val) {
    AppService.SetAutoStart().then(() => {
      message.success("设置成功")
    })
  } else {
    AppService.RemoveAutoStart().then(() => {
      message.success("移除成功")
    })
  }
}
</script>

<style scoped>

</style>