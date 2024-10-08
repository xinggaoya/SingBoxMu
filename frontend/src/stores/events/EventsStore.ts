import {defineStore} from 'pinia'
import {ref} from 'vue'
import {GetLogs, GetMemory, GetTraffic} from "@/api/clash/ClashApi";

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
            // Events.On("logs", (data: any) => {
            //     logs.value.push(JSON.parse(data.data))
            //     // 仅保留最新100条
            //     if (logs.value.length > 50) {
            //         logs.value.shift()
            //     }
            // })
            // Events.On("traffic", (data: any) => {
            //     const info = JSON.parse(data.data)
            //     // 转换为MB 保留小数点后2位
            //     info.up = (info.up / 1024 / 1024)
            //     info.down = (info.down / 1024 / 1024)
            //     traffic.value.up = info.up
            //     traffic.value.down = info.down
            //     totalUseTraffic.value += (info.up + info.down)
            // })
            // Events.On("memory", (data: any) => {
            //     const info = JSON.parse(data.data)
            //     // 转换为MB 保留小数点后2位
            //     info.inuse = (info.inuse / 1024 / 1024)
            //     info.oslimit = (info.oslimit / 1024 / 1024)
            //     memory.value = info
            // })

            GetLogs((_, chunk) => {
                const text = new TextDecoder("utf-8").decode(chunk);
                if (text) {
                    logs.value.push(JSON.parse(text))
                    // 仅保留最新100条
                    if (logs.value.length > 50) {
                        logs.value.shift()
                    }
                }
            })
            GetMemory((_, chunk) => {
                const text = new TextDecoder("utf-8").decode(chunk);
                if (text) {
                    const info = JSON.parse(text)
                    // 转换为MB 保留小数点后2位
                    info.inuse = (info.inuse / 1024 / 1024)
                    info.oslimit = (info.oslimit / 1024 / 1024)
                    memory.value = info
                }
            })
            GetTraffic((_, chunk) => {
                const text = new TextDecoder("utf-8").decode(chunk);
                if (text) {
                    const info = JSON.parse(text)
                    // 转换为MB 保留小数点后2位
                    info.up = (info.up / 1024 / 1024)
                    info.down = (info.down / 1024 / 1024)
                    traffic.value.up = info.up
                    traffic.value.down = info.down
                    totalUseTraffic.value += (info.up + info.down)
                }
            })

        }

        return {logs, traffic, memory, listen, totalUseTraffic}
    },
    {
        persist: {
            storage: localStorage,
            paths: ['totalUseTraffic']
        }
    }
)
