import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  server: {
    proxy: {
      // 当请求路径以 `/api` 开头时，代理请求到目标地址
      '/api': {
        target: 'http://localhost:8080',  // 目标服务器的地址
        changeOrigin: true,                // 如果是跨域请求，修改请求头中的 `Origin` 前缀
        logLevel: 'debug',
      },
    },
  },

  plugins: [
    vue(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})


