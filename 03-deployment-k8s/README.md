# Deployment

## 扩容
通过`Deployment`资源可以管理`Pod`进行扩容

## Rolling Update
通过`滚动更新`来确保新版本未`ready`之前，先不删除旧版本的`Pod`
```bash
# 版本回滚和历史版本查看
kubectl rollout undo deployment hellok8s-deployment
kubectl rollout history deployment hellok8s-deployment
kubectl rollout undo deployment/hellok8s-deployment --to-revision=2
```
![](https://camo.githubusercontent.com/526eeefe7af67cd567486657693e9acac102f726f5b69427b3796bd0c3ff17be/68747470733a2f2f63646e2e6a7364656c6976722e6e65742f67682f6775616e677a68656e676c692f50696355524c406d61737465722f755069632f726f6c6c696e677570646174652e706e67)

## livenessProb
通过`存活探针`来判断`容器`是否正常运行`容器`异常时`kubelet`使用`存活探针`来判断什么时候重启容器

## readiness
> 就绪探测器可以知道容器何时准备好接受请求流量，当一个 Pod 内的所有容器都就绪时，才能认为该 Pod 就绪。 这种信号的一个用途就是控制哪个 Pod 作为 Service 的后端。 若 Pod 尚未就绪，会被从 Service 的负载均衡器中剔除。-- ReadinessProb

在生产环境中，升级服务的版本是日常的需求，这时我们需要考虑一种场景，即当发布的版本存在问题，就不应该让它升级成功。kubelet 使用就绪探测器可以知道容器何时准备好接受请求流量，当一个 pod 升级后不能就绪，即不应该让流量进入该 pod，在配合 rollingUpate 的功能下，也不能允许升级版本继续下去，否则服务会出现全部升级完成，导致所有服务均不可用的情况。