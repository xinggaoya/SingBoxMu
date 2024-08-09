import {defineStore} from 'pinia'
import {ref} from 'vue'

interface SettingsConfig {
    autoStart: boolean
    autoRun: boolean
}


export const useConfigStore = defineStore('config', () => {

        const form = ref<SettingsConfig>({
            autoStart: false,
            autoRun: false
        })

        return {form}
    },
    {persist: true}
)
