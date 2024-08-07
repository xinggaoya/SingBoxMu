import './assets/main.scss'
// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'

import {createApp} from 'vue'
import {usePinia} from '@/stores/index.ts'

import App from './App.vue'
import router from './router'
import {createDirectives} from "@/directives";

const app = createApp(App)

usePinia(app)
app.use(router)
createDirectives(app)

app.mount('#app')
