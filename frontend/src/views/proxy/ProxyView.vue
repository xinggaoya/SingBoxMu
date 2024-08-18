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
import {useMessage} from "naive-ui";
import {GetProxies, SwitchProxy} from "@/api/clash/ClashApi";

const proxies = ref<any>({})
const message = useMessage()

onMounted(() => {
  GetProxies().then((res: any) => {
    // 循环对象属性
    Object.keys(res.proxies).forEach(key => {
      // 获取属性值
      const value = res.proxies[key]
      if (value.all?.length > 0) {
        proxies.value[key] = value
      }
    })
  })
})

function switchProxy(group: any, name: string) {
  SwitchProxy(group, name).then(() => {
    message.success("切换成功")
  }).catch(()=>{
    message.error("切换失败")
  })
}

</script>

<style scoped>
.mr-2 {
  cursor: pointer;
}

</style>