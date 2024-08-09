import {defineStore} from 'pinia'
import {ref} from 'vue'


export const useAppStore = defineStore('app', () => {
    // 运行状态
    const isRunning = ref(false)
    // 代理模式
    const proxyMode = ref('')
    const isDark = ref(false)

    function toggleDark() {
        isDark.value = !isDark.value
    }

    return {
        isRunning,
        proxyMode,
        isDark,
        toggleDark,
    }
}, {
    persist: true,
})

export default useAppStore