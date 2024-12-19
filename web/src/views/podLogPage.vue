
<template>
  <div class="container">
  <DynamicScroller
      :items="messages"
      :min-item-size="10"
      class="scroller"
      :max-visible-items="200"
  >
    <template #default="{ item, index, active }">
      <DynamicScrollerItem class="message"
          :item="item"
          :active="active"
          :data-active="active"
          :size-dependencies="[
          item,
        ]"
          :data-index="index"
      >

        <div class="text" v-html="item.message"></div>
      </DynamicScrollerItem>
    </template>
  </DynamicScroller>
    </div>
</template>





<script>
import { AnsiUp } from 'ansi_up';

export default {
  props: {
    messages: {
      type: Array,
      default: () => [],
    },
  },

  name: "podLogPage",
  data() {
    return {
      messageBuffer: [],
      socket: null, // WebSocket 实例
      messages: [],
      itemSize: 100,
      id: 0,
      minItemSize: 5,
    };
  },
  methods: {

    handleWebSocketMessage(data) {
      const messageWithId = {
        id: Date.now(),  // 使用时间戳作为唯一ID
        message: data,    // 原始消息内容
      };


      this.messageBuffer.push(messageWithId);  // 将新消息推入缓冲区

      if (this.bufferTimeout) {
        clearTimeout(this.bufferTimeout);  // 清除之前的定时器
      }
      this.bufferTimeout = setTimeout(() => {
        this.messages.push(...this.messageBuffer);  // 将缓冲区的消息批量推送到 messages
        this.messageBuffer = [];  // 清空缓冲区
      }, 50);
    },

    connectWebSocket() {
      const currentUrl = window.location.href.split('?')[0];
      this.socket = new WebSocket("ws://localhost:8080/api/getpodlogs")
      const urlParams = new URLSearchParams(window.location.search);
      const namespace = urlParams.get('nameSpace');
      const podname = urlParams.get('podName');

      const ansi_up = new AnsiUp();
      // 监听消息事件
      this.socket.onmessage = (event) => {
        const  data = JSON.parse(JSON.stringify(event.data));
        this.handleWebSocketMessage(ansi_up.ansi_to_html(data));
      };


      // 监听连接打开事件
      this.socket.onopen = () => {
        console.log('WebSocket connected');
        const authMessage = JSON.stringify({
          namespace: namespace,
          podname:  podname,
        });
        console.log(authMessage);
        this.socket.send(authMessage); // 向服务器发送消息
      };

      // 监听连接关闭事件
      this.socket.onclose = () => {
        console.log('WebSocket disconnected');
      };

      // 监听错误事件
      this.socket.onerror = (error) => {
        console.error('WebSocket error', error);
      };
    },

    beforeUnmount() {
      // 在组件销毁前关闭 WebSocket 连接
      if (this.socket) {
        this.socket.close();
      }
    },
    },
  mounted() {
    this.connectWebSocket();
  }
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100vh; /* 父容器高度 */
}

.scroller {
  flex: 1; /* 让 scroller 充满父容器 */
  overflow-y: auto; /* 设置 overflow-y 为 auto */
  border: solid 1px #42b983;
}

.message{
  display: flex;
  min-height: 32px;
  padding: 12px;
  box-sizing: border-box;
}
</style>