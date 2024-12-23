import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css';
import {createPinia} from "pinia";
import piniaPersistedstate from 'pinia-plugin-persistedstate'
import VueVirtualScroller from 'vue-virtual-scroller'
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const app = createApp(App)
const pinia = createPinia()
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
pinia.use(piniaPersistedstate)
app.use(pinia)
app.use(ElementPlus)
app.use(VueVirtualScroller)
app.use(router)
app.mount('#app')
