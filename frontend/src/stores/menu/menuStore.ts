import {defineStore} from 'pinia'
import {ref} from 'vue'

export interface MenuStore {
    label: string
    key: string
    icon?: any | null
    children?: MenuStore[]
}

export const useMenuStore = defineStore('menu', () => {
    const menuList = ref<MenuStore[]>([])
    // 当前选中的菜单
    const currentMenu = ref<number>(0)

    const roleName = ref<string>('')

    const list = ref<any>([])

    // 初始化菜单
    async function initMenu() {
    }

    // 根据key获取数据
    function setRoleNameByKey(key: number) {
        roleName.value = list.value.find((item: any) => item.id === key)?.modelRole.title
    }

    return {menuList, roleName, currentMenu, initMenu, setRoleNameByKey}
})
