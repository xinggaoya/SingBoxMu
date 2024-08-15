<template>
  <n-card style="height: 100%" content-style="height: 100%">
    <n-scrollbar>
      <n-flex>
        <n-collapse>
          <n-collapse-item
              v-for="(value,key,index) in proxies"
              :key="index"
              :title="key"
              :name="index">
            <n-flex>
              <n-tag type="primary"
                     class="mr-2"
                     v-for="(item,i) in value.all"
                     :key="i"
                     @click="switchProxy(key, item)">
                {{ item }}
              </n-tag>
            </n-flex>
          </n-collapse-item>
        </n-collapse>

      </n-flex>
    </n-scrollbar>
  </n-card>
</template>
<script setup lang="ts">
import {onMounted, ref} from "vue";
import {ClashService} from "@api/changeme/app/service";
import {useMessage} from "naive-ui";

const proxies = ref<any>({})
const message = useMessage()

onMounted(() => {
  ClashService.GetProxies().then(res => {
    if (res.code === 10000) {
      const info = JSON.parse(res.data)
      console.log(info)
      // 循环对象属性
      Object.keys(info.proxies).forEach(key => {
        // 获取属性值
        const value = info.proxies[key]
        if (value.all.length > 0) {
          proxies.value[key] = value
        }
      })
    }
  })
})

function switchProxy(group: any, name: string) {
  ClashService.SwitchProxy(group, name).then(res => {
    if (res.code === 10000) {
      message.success("切换成功")
    } else {
      message.error(res.msg)
    }
  })
}

</script>

<style scoped>
.mr-2 {
  cursor: pointer;
}

</style>