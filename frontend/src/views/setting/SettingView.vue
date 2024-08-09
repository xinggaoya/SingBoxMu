<template>
  <n-card title="Setting" style="height: 100%" content-style="height: 100%;">
    <n-scrollbar>
      <n-divider>系统</n-divider>
      <n-form :model="configStore.form" ref="formRef" label-placement="left">
        <n-form-item label="开机自启" path="autoStart">
          <n-switch v-model:value="configStore.form.autoStart" @update-value="updateAutoStart"/>
        </n-form-item>
      </n-form>
    </n-scrollbar>
  </n-card>
</template>

<script setup lang="ts">
import {useConfigStore} from "@/stores/config/ConfigStore.ts";
import {AppService} from "@api/changeme/app/service";
import {onMounted, ref} from "vue";
import {useMessage} from "naive-ui";

const configStore = useConfigStore()
const formRef = ref()
const message = useMessage()

onMounted(() => {
  AppService.GetVersion().then((res: any) => {
    console.log(res.data)
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