# Namespace

在实际的开发当中，有时候我们需要不同的环境来做开发和测试，例如 `dev` 环境给开发使用，`test` 环境给 QA 使用，那么 k8s 能不能在不同环境 `dev` `test` `uat` `prod` 中区分资源，让不同环境的资源独立互相不影响呢，答案是肯定的，k8s 提供了名为 Namespace 的资源来帮助隔离资源。

在 Kubernetes 中，名字空间（Namespace） 提供一种机制，将同一集群中的资源划分为相互隔离的组。 同一名字空间内的资源名称要唯一，但跨名字空间时没有这个要求。 名字空间作用域仅针对带有名字空间的对象，例如 Deployment、Service 等。

``` bash
# -n namespace 指定命名空间
kubectl apply -f deployment.yaml -n dev

kubectl get pods -n dev
```