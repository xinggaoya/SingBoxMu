import {defineStore} from 'pinia'
import {ref} from 'vue'
import {useRouter} from "vue-router";


export const useAppStore = defineStore('app', () => {
    const isDark = ref(false)
    const userToken = ref('')
    const router = useRouter()

    function toggleDark() {
        isDark.value = !isDark.value
    }

    function getUserToken() {
        return userToken.value
    }

    function setToken(token: string) {
        userToken.value = token
    }

    function clearToken() {
        userToken.value = ''
        router.push('/login')
    }

    return {
        isDark,
        userToken,
        toggleDark,
        setToken,
        clearToken,
        getUserToken
    }
}, {
    persist: true,
})

export default useAppStore