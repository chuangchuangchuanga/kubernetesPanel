<template>
  <div ref="scroller">
    <DynamicScroller :items="messages"  :item-size="itemSize" >
      <template #default="{ item, index }">
        <div class="item" >
          <h3>{{ index }}</h3>
          <p >{{ item }}</p>
          <el-divider />
        </div>
      </template>
    </DynamicScroller>
  </div>
</template>


<script>


export default {



  name: "podLogPage",
  data() {
    return {
      socket: null, // WebSocket 实例
      messages: [],
      itemSize: 50,
    };
  },
  methods: {
    scrollToBottom() {
      this.$nextTick(() => {
        window.scrollTo(0, document.body.scrollHeight);
      });
    },

    connectWebSocket() {
      const currentUrl = window.location.href.split('?')[0];
      this.socket = new WebSocket("ws://localhost:8080/api/getpodlogs")
      const urlParams = new URLSearchParams(window.location.search);
      const namespace = urlParams.get('nameSpace');
      const podname = urlParams.get('podName');


      // 监听消息事件
      this.socket.onmessage = (event) => {
        console.log(event.data)
        const  data = JSON.parse(JSON.stringify(event.data));
        this.messages.push(data)
        this.scrollToBottom()
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


