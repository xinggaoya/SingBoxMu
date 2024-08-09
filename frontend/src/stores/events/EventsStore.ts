import {defineStore} from 'pinia'
import {ref} from 'vue'
import {Events} from "@wailsio/runtime";

interface logConfig {
    type: string,
    payload: string
}

interface TrafficConfig {
    up: number,
    down: number
}

interface MemoryConfig {
    inuse: number,
    oslimit: number
}


export const useEventsStore = defineStore('Events', () => {

        // 日志
        const logs = ref<logConfig[]>([])
        // 流量
        const traffic = ref<TrafficConfig>({down: 0, up: 0})
        // 内存
        const memory = ref<MemoryConfig>({inuse: 0, oslimit: 0})

        // 流量统计
        const totalUseTraffic = ref<number>(0)


        // 监听事件
        function listen() {
            Events.On("logs", (data: any) => {
                logs.value.push(JSON.parse(data.data))
                // 仅保留最新100条
                if (logs.value.length > 50) {
                    logs.value.shift()
                }
            })
            Events.On("traffic", (data: any) => {
                const info = JSON.parse(data.data)
                traffic.value = info
                totalUseTraffic.value = info.up + info.down
            })
            Events.On("memory", (data: any) => {
                const info = JSON.parse(data.data)
                // 转换为MB
                info.inuse = info.inuse / 1024 / 1024
                info.oslimit = info.oslimit / 1024 / 1024
                memory.value = info
            })
        }

        return {logs, traffic, memory, listen, totalUseTraffic}
    },
    {persist: true}
)
