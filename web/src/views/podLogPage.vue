
<template>
  <div class="container">
  <DynamicScroller
      :items="messages"
      :min-item-size="20"
      class="scroller"
      :max-visible-items="100"
  >
    <template #default="{ item, index, active }">
      <DynamicScrollerItem class="message"
          :item="item"
          :active="active"
          :data-active="active"
          :size-dependencies="[
          item,
        ]"
          :data-index="item.id"
      >
        <div class="message" v-html="item.message" ></div>
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

      this.messages.push(messageWithId);

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
  background-color: RGB(67,67,67);
}

.scroller {
  flex:  1; /* 让 scroller 充满父容器 */
  overflow-y: auto; /* 设置 overflow-y 为 auto */
  border: solid 1px #42b983;
}

.message  span {
  color: #ffffff !important;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}
</style>