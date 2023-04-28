# Secret
上面提到，我们会选择以 configmap 的方式挂载配置信息，但是当我们的配置信息需要加密的时候， configmap 就无法满足这个要求。例如上面要挂载数据库密码的时候，就需要明文挂载。

这个时候就需要 Secret 来存储加密信息，虽然在资源文件的编码上，只是通过 Base64 的方式简单编码，但是在实际生产过程中，可以通过 pipeline 或者专业的 AWS KMS 服务进行密钥管理。这样就大大减少了安全事故。

Secret 的资源定义和 ConfigMap 结构基本一致，唯一区别在于 kind 是 `Secret`，还有 Value 需要 Base64 编码，你可以通过下面命令快速 Base64 编解码。当然 Secret 也提供了一种 `stringData`，可以不需要 Base64 编码。
```bash
echo "db_password" | base64
# ZGJfcGFzc3dvcmQK

echo "ZGJfcGFzc3dvcmQK" | base64 -d
# db_password
```