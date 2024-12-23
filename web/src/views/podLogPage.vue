
<template>
  <div class="container">
  <DynamicScroller
      :items="messages"
      :min-item-size="20"
      class="scroller"
      id="scroller"
  >
    <template #default="{ item, index, active }">
      <DynamicScrollerItem
          class="scroller-item"
          :item="item"
          :data-active="active"
          :active="active"
          :data-index="item.id"
          :size-dependencies="[
            item.message,
          ]">
        <div class="message"  v-html="item.message"></div>

      </DynamicScrollerItem>
    </template>
  </DynamicScroller>


    <el-button class="floating-button" type="primary" @click="changeAutoScroll">

      <el-icon v-if="autoScroll===false">
        <Lock />
      </el-icon>
      <el-icon v-else-if="autoScroll===true"><Unlock /></el-icon>
      滚动
    </el-button>
  </div>


</template>



<script>
import { AnsiUp } from 'ansi_up';

export default {
  name: "podLogPage",
  props: {
    active: {
      type: Boolean,
      required: true
    },
  },
  data() {
    return {
      socket: null, // WebSocket 实例
      messages: [],
      id: 0,
      autoScroll: true,
    };
  },
  methods: {
    changeAutoScroll() {
      this.autoScroll = !this.autoScroll;
      if (this.autoScroll) {
        this.scrollToBottom()
      }
    },

    handleWebSocketMessage(data) {
      const messageWithId = {
        id: Date.now(),  // 使用时间戳作为唯一ID
        message: data,    // 原始消息内容
      };
      this.messages.push(messageWithId);
      if (this.messages.length >= 1000) {
        // 如果数组长度大于等于1000，移除最旧的一条消息
        this.messages.shift();
      }
      if (this.autoScroll===true) {
        this.scrollToBottom()
      }

    },


    scrollToBottom(){
      var scrollableDiv = document.getElementById('scroller');
      // 将元素滚动到底部
      scrollableDiv.scrollTop = scrollableDiv.scrollHeight;
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
  overflow-y: auto;
}

.scroller {
  flex: auto 1 1 ;
  overflow-y: auto; /* 设置 overflow-y 为 auto */
  border: solid 1px #42b983;
}

.floating-button {
  position: fixed;
  right: 20px;
  bottom: 80px;
  z-index: 9999;
}

</style>