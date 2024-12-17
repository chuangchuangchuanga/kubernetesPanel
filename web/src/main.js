import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css';
import {createPinia} from "pinia";
import piniaPersistedstate from 'pinia-plugin-persistedstate'
import VueVirtualScroller from 'vue-virtual-scroller'

const app = createApp(App)
const pinia = createPinia()


pinia.use(piniaPersistedstate)
app.use(pinia)
app.use(ElementPlus)
app.use(VueVirtualScroller)
app.use(router)
app.mount('#app')
