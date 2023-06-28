# ubuntu部署单节点Kubernetes1.27

系统：Ubuntu22.04

### 1.准备工作

```shell
curl -fsSL https://get.docker.com | sudo sh
# 使用 aliyun 的 k8s 源安装 kubeadm 和相关命令行工具
apt-get update && apt-get install -y apt-transport-https
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -
echo "deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main" > /etc/apt/sources.list.d/kubernetes.list
apt-get update
apt-get install -y kubelet kubeadm kubectl
systemctl enable kubelet
# 修改 docker 的 cgroup driver 选项为 systemd，与 k8s 保持一致，并修改 registry-mirror 加速下载
vi /etc/docker/daemon.json
    {
        "exec-opts": ["native.cgroupdriver=systemd"],
        "registry-mirrors": ["https://registry.docker-cn.com", "https://docker.mirrors.ustc.edu.cn"]
    }
systemctl restart docker
```

### 2.使用 kubeadm init k8s

```shell
# 使用阿里云上提供的 k8s 镜像（这里指定的网络与后续使用的网络插件的配置保持一致）
kubeadm init --image-repository registry.aliyuncs.com/google_containers \
    --service-cidr=10.1.0.0/16 --pod-network-cidr=10.244.0.0/16
```

#### 2.1 问题修复（如果上一步执行正常无报错 则跳过）

```shell
container runtime is not running :CRI v1 runtime API is not implemented for endpoint \"unix:///var/run/containerd/containerd.sock\": rpc error: code = Unimplemented desc = unknown service runtime.v1.RuntimeService
```

报错的原因是ubuntu22.04默认预装的是旧版本的containerd.io 需要下载二进制包手动替换：

以下是简单步骤（可以照做）具体详细的文档参见https://github.com/containerd/containerd/blob/main/docs/getting-started.md

```shell
#下载containerd二进制包
wget https://github.com/containerd/containerd/releases/download/v1.7.2/containerd-1.7.2-linux-amd64.tar.gz
#将其解压缩到/usr/local下:
tar Cxzvf /usr/local containerd-1.7.2-linux-amd64.tar.gz
#接下来从runc的github上单独下载安装runc，该二进制文件是静态构建的，并且应该适用于任何Linux发行版。
wget https://github.com/opencontainers/runc/releases/download/v1.1.7/runc.amd64
install -m 755 runc.amd64 /usr/local/sbin/runc
#生成containerd的配置文件
mkdir -p /etc/containerd
containerd config default > /etc/containerd/config.toml
```

根据官方文档指导 需要将`systemd` 设置为 cgroup 驱动对于使用systemd作为init system的Linux的发行版，使用systemd作为容器的cgroup driver可以确保服务器节点在资源紧张的情况更加稳定 详见https://kubernetes.io/zh-cn/docs/setup/production-environment/container-runtimes/

修改刚刚生成的配置文件

```shell
vim /etc/containerd/config.toml
```

```toml
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
  ...
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true
```

国内环境需要替换k8s官方镜像仓库为阿里云仓库

还是刚刚的配置文件：

```toml
[plugins."io.containerd.grpc.v1.cri"]
  ...
  # sandbox_image = "registry.k8s.io/pause:3.8"
  sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.9"
```

为了通过systemd启动containerd，请还需要从`https://raw.githubusercontent.com/containerd/containerd/main/containerd.service`下载`containerd.service`单元文件，并将其放置在` /etc/systemd/system/containerd.service`中。 配置containerd开机启动，并启动containerd，执行以下命令:

```shell
cd /etc/systemd/system
wget https://raw.githubusercontent.com/containerd/containerd/main/containerd.service
systemctl daemon-reload
systemctl enable containerd --now 
```

下载安装crictl工具（选装，ubuntu22.04自带）:

```shell
wget https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.27.0/crictl-v1.27.0-linux-amd64.tar.gz
tar -zxvf crictl-v1.27.0-linux-amd64.tar.gz
install -m 755 crictl /usr/local/bin/crictl
```

使用crictl测试一下，确保可以打印出版本信息并且没有错误信息输出:

```shell
crictl --runtime-endpoint=unix:///run/containerd/containerd.sock  version

Version:  0.1.0
RuntimeName:  containerd
RuntimeVersion:  v1.7.2
RuntimeApiVersion:  v1
```

notice：这里是第二个坑，执行命令返回 

```shell
crictl --runtime-endpoint=unix:///run/containerd/containerd.sock  version
FATA[0000] validate service connection: CRI v1 runtime API is not implemented for endpoint "unix:///run/containerd/containerd.sock": rpc error: code = Unimplemented desc = unknown service runtime.v1.RuntimeService
```

！！（上一步没问题的跳过）原因是系统预装的containerd还在运行，需要删除掉后重启 

```shell
apt remove containerd.io
systemctl daemon-reload
systemctl enable containerd --now
systemctl status containerd
```

此时再次运行检查命令，返回正常

```shell
crictl --runtime-endpoint=unix:///run/containerd/containerd.sock  version
```

启动kubelet

```shell
systemctl enable kubelet.service
```

再次运行

```shell
 kubeadm init --image-repository registry.aliyuncs.com/google_containers \
    --service-cidr=10.1.0.0/16 --pod-network-cidr=10.244.0.0/16
```

#### 2.2 完成安装

全部执行完成后根据提示执行下述命令

```shell
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:
	
  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

#单节点集群跳过这步
kubeadm join 172.16.3.169:6443 --token z4o3n1.l64vso6rqo8pobmp \
	--discovery-token-ca-cert-hash sha256:124049c1f0335b2ed1314262710e60ac4d344403bd6375d2c66ea7e3013da99c
```

执行 kubectl get pods -A

```shell
NAMESPACE     NAME                                  READY   STATUS    RESTARTS   AGE
kube-system   coredns-7bdc4cb885-w6vrn              0/1     Pending   0          5s
kube-system   coredns-7bdc4cb885-wx2nr              0/1     Pending   0          5s
kube-system   etcd-***-****-**                      1/1     Running   0          5s
kube-system   kube-apiserver-***-****-**            1/1     Running   0          5s
kube-system   kube-controller-manager-***-****-**   1/1     Running   0          5s
kube-system   kube-proxy-lbcbk                      1/1     Running   0          5s
kube-system   kube-scheduler-***-****-**            1/1     Running   0          5s
```

大功告成

### 3.安装包管理器helm 3

Helm是Kubernetes的包管理器，后续流程也将使用Helm安装Kubernetes的常用组件。 这里先在master节点node1上安装helm。

```shell
wget https://get.helm.sh/helm-v3.12.0-linux-amd64.tar.gz
tar -zxvf helm-v3.12.0-linux-amd64.tar.gz
mv linux-amd64/helm  /usr/local/bin/
```

执行 helm list 确认没有错误输出