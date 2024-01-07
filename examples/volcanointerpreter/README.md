1. 启动本地开发集群的时候
```shell
sudo sysctl fs.inotify.max_user_instances=1280
sudo sysctl fs.inotify.max_user_watches=655360

export CHINA_MAINLAND=1
```

2. 部署集群的时候修改hack/deploy-karmada.sh中的地址参数，使其等于webhook的ip地址
```shell
interpreter_webhook_example_service_external_ip_address=${interpreter_webhook_example_service_external_ip_prefix}.8
```

3. 获取证书，替换webhook-configurations中的caBundle，同时修改url的值为webhook的地址和端口
```shell
cat ${HOME}/.karmada/ca.crt | base64 | tr "\n" " "|sed s/[[:space:]]//g
```


4. 生成webhook server的TLS证书
```shell
kubectl get secret webhook-cert -n karmada-system -o jsonpath="{.data.tls\.crt}" | base64 -d > tls.crt
kubectl get secret webhook-cert -n karmada-system -o jsonpath="{.data.tls\.key}" | base64 -d > tls.key
```

5. 添加如下启动参数
```shell
--kubeconfig=/root/.kube/karmada.config
--bind-address=0.0.0.0
--secure-port=8445
--cert-dir=/root/karmada/_output/bin/linux/amd64
--v=4
```