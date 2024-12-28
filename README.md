# KubernetesPanel介绍
KubernetesPanel 是一个基于 Kubernetes 的监控面板，用于实时监控 Kubernetes 集群中的资源使用情况、容器状态、节点状态等。

## 快速开始
```kubernetes apply -f https://raw.githubusercontent.com/chuangchuangchuanga/kubernetesPanel/refs/heads/main/deploy/one-in-file.yaml```

要访问KubernetesPanel，需要添加Ingress ,根据自己情况添加Ingress信息，访问地址需要使用使用 /kubernetes 前缀
```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: test-k8spanel-ingress
spec:
  rules:
    - host: example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: k8s-panel-svc
                port:
                  number: 8080
```


## 特点
- 查看所有Deployment
- 查看所有的Pod状态
- 查看Pod前台页面输出的日志(使用了vue-virtual-scroller虚拟化的列表，可以在大量数据下提供高性能的渲染和滚动，减少浏览器内存占用)




