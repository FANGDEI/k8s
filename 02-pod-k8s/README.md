# Pod
```yaml
apiVersion: v1
kind: Pod # 创建的资源类型
metadata:
  name: nginx-pod # metadata.name 创建的 Pod 名字, 名字唯一
spec:
  containers: # spec.containers 要运行的容器名称和镜像名称 镜像默认来源 DockerHub
    - name: nginx-container
      image: nginx
```
```bash
kubectl apply -f nginx.yaml
# pod/nginx-pod created

kubectl get pods
# nginx-pod         1/1     Running   0           6s

kubectl port-forward nginx-pod 4000:80
# Forwarding from 127.0.0.1:4000 -> 80
# Forwarding from [::1]:4000 -> 80
```
`kubectl exec -it` 可以用来进入 `Pod` 内容器的 `Shell`。通过命令下面的命令来配置 `nginx` 的首页内容。

```bash
kubectl exec -it nginx-pod /bin/bash # 过时
kubectl exec -it nginx-pod -- /bin/bash

echo "hello kubernetes by nginx!" > /usr/share/nginx/html/index.html

kubectl port-forward nginx-pod 4000:80
```
最后可以通过浏览器或者 `curl` 来访问 `http://127.0.0.1:4000` , 查看是否成功启动 `nginx` 和返回字符串 `hello kubernetes by nginx!`。

# Pod 与 Container 的不同
回到 `pod` 和 `container` 的区别，我们会发现刚刚创建出来的资源如下图所示，在最内层是我们的服务 `nginx`，运行在 `container` 容器当中， `container` (容器) 的本质是进程，而 `pod` 是管理这一组进程的资源。

![](https://camo.githubusercontent.com/5cf6cf3d7535429968e20836e2e5312bef344257fa62647ba4928129cd4f40a4/68747470733a2f2f63646e2e6a7364656c6976722e6e65742f67682f6775616e677a68656e676c692f50696355524c406d61737465722f755069632f6e67696e785f706f642e706e67)

所以自然 `pod` 可以管理多个 `container`，在某些场景例如服务之间需要文件交换(日志收集)，本地网络通信需求(使用 localhost 或者 Socket 文件进行本地通信)，在这些场景中使用 `pod` 管理多个 `container` 就非常的推荐。而这，也是 k8s 如何处理服务之间复杂关系的第一个例子，如下图所示：
![](https://camo.githubusercontent.com/0c5ac31305f9d3a9bea50e0211b59d3fd033ff9fb5a5b0d70d75bf5dc9ca84c0/68747470733a2f2f63646e2e6a7364656c6976722e6e65742f67682f6775616e677a68656e676c692f50696355524c406d61737465722f755069632f706f642e706e67)
# Pod 其它命令
```bash
kubectl logs --follow nginx-pod
                              
kubectl exec nginx-pod -- ls

kubectl delete pod nginx-pod
# pod "nginx-pod" deleted

kubectl delete -f nginx.yaml
# pod "nginx-pod" deleted
```
