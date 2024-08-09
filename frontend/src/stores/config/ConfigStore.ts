import {defineStore} from 'pinia'
import {ref} from 'vue'

interface SettingsConfig {
    autoStart: boolean
}


export const useConfigStore = defineStore('config', () => {

        const form = ref<SettingsConfig>({
            autoStart: false
        })

        return {form}
    },
    {persist: true}
)
