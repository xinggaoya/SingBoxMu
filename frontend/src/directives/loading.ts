import { App, createVNode, render, VNode } from 'vue'
import Loading from '@/components/loading/index.vue'

const vnode: VNode = createVNode(Loading) as VNode

// 页面加载指令
export function loading(app: App) {
  app.directive('loading', {
    mounted: (el: HTMLElement, binding: { value: any }) => {
      const { value } = binding
      renderLoading(value, el)
    },
    updated: (el: HTMLElement, binding: { value: any }) => {
      const { value } = binding
      renderLoading(value, el)
    },
    unmounted: (el: HTMLElement) => {
      renderLoading(false, el)
    },
  })
}

function renderLoading(value: boolean, el: HTMLElement) {
  el.style.position = 'relative'
  if (value) {
    render(vnode, el)
  } else {
    render(null, el)
  }
}
