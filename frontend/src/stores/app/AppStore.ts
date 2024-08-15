import {defineStore} from 'pinia'
import {ref} from 'vue'
import {ClashService} from "@api/changeme/app/service";
import {useEventsStore} from "@/stores/events/EventsStore";


interface kernelVersionConfig {
    meta: string
    premium: string
    version: string
}

export const useAppStore = defineStore('app', () => {
    // 运行状态
    const isRunning = ref<boolean>(false)
    // 运行模式
    const isAdminRun = ref<boolean>(false)
    // 内核版本
    const kernelVersion = ref<kernelVersionConfig>({
        meta: '',
        premium: '',
        version: ''
    })
    // 代理模式
    const proxyMode = ref('')
    // 暗黑模式
    const isDark = ref(false)

    const events = useEventsStore()

    // 获取版本
    function getKernelVersion():Promise<void> {
        return new Promise((resolve) => {
            ClashService.GetVersion().then((res) => {
                if (res.code === 10000) {
                    isRunning.value = true
                    kernelVersion.value = JSON.parse(res.data)
                    events.listen()
                } else {
                    isRunning.value = false
                }
                resolve()
            })
        })
    }

    function toggleDark() {
        isDark.value = !isDark.value
    }

    return {
        isRunning,
        proxyMode,
        isDark,
        toggleDark,
        kernelVersion,
        getKernelVersion,
        isAdminRun,
    }
}, {
    persist: true,
})

export default useAppStore