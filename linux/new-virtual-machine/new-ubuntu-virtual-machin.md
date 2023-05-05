# 轻巧创建全新的Ubuntu 虚拟机

## 1. 获取 Ubuntu 官网提供的虚拟机镜像（已安装好的磁盘镜像）

1.1. 浏览页面 https://cloud-images.ubuntu.com/releases/bionic/release/ 查找 `-server-cloudimg-amd64.img   ` 后缀的最新镜像，复制其下载地址。（目前是：https://cloud-images.ubuntu.com/releases/bionic/release/ubuntu-18.04-server-cloudimg-amd64.img）

1.2. 进入放置虚拟机镜像的目录，后面的操作都会基于这个目录

```sh
cd /data
```

注意：不一定要是`/data`，任意目录都可以，根据当前机器上磁盘分区的使用情况，找到合适的目录即可。

1.3. 下载虚拟机镜像到当前目录

```sh
curl -k -LO https://cloud-images.ubuntu.com/releases/bionic/release/ubuntu-18.04-server-cloudimg-amd64.img
```

下载完成后可以用 ls 命令确认 `ubuntu-18.04-server-cloudimg-amd64.img` 文件存在

```sh
ls
# 输出：
# ubuntu-18.04-server-cloudimg-amd64.img
```

1.4. 重命名文件名为 `ubuntu-18-{姓名拼音}.qcow2`，带上使用者的姓名拼音是为了方便后期维护管理。

```sh
# 设置 MY_NAME 环境变量，注意这个环境变量后面会一直用到
MY_NAME=myname

mv  ubuntu-18.04-server-cloudimg-amd64.img  ubuntu-18-${MY_NAME}.qcow2
```

## 2. 定制虚拟机镜像

2.1. 首先需要给虚拟机内的磁盘扩容，需要使用 libguestfs 工具链，但是本机不一定有安装，这里选择使用 `bkahlert/libguestfs` docker 镜像，把本地镜像挂入容器内执行命令：

```sh
docker run -it --rm \
  -v "${PWD}:/data" \
  -w /data \
  -u "$(id -u):$(id -g)" \
  --entrypoint /usr/bin/env \
  bkahlert/libguestfs \
  -- qemu-img resize ubuntu-18-${MY_NAME}.qcow2 100G
```

上面的命令只是把虚拟机内部的磁盘空间设置为100G，`ubuntu-18-${MY_NAME}.qcow2` 的大小不变，大概为300M~400M，根据使用 `ubuntu-18-${MY_NAME}.qcow2` 文件会自动增长，但是这样性能不好（文件的对应磁盘空间可能不连续），可以预先把文件的大小也拓展到100G

```
dd if=/dev/zero bs=1G seek=100 count=0 of=ubuntu-18-${MY_NAME}.qcow2
```

确认拓展后的文件大小为100G：

```
ls -alh ubuntu-18-${MY_NAME}.qcow2
```

2.2. 需要更改虚拟机镜像的用户密码（默认用户名为ubuntu），当前使用的 Ubuntu Could 镜像比较特殊，需要定制一个包含用户密码的iso镜像，然后让这个iso镜像**每次**都随虚拟机启动，实现修改用户密码的效果。

需要先安装 cloud-image-utils 包：

```sh
apt install cloud-image-utils
```

然后定制 cloud-config 文件，创建文件 `cloud-config-${MY_NAME}.txt`：

```sh
cat >cloud-config-${MY_NAME}.txt <<EOF
#cloud-config
password: 12345
chpasswd: { expire: False }
ssh_pwauth: True
hostname: localhost
EOF
```

注意：上面的12345就是用户密码

最终生成 `cloud-config-${MY_NAME}.iso` 文件

```sh
cloud-localds cloud-config-${MY_NAME}.iso cloud-config-${MY_NAME}.txt

# cloud-config 可以删除了
rm cloud-config-${MY_NAME}.txt
```

## 3. 运行虚拟机

3.1. 创建并且启动虚拟机

```sh
virt-install \
  --name ubuntu-18-${MY_NAME} \
  --os-type=Linux \
  --vcpus 8 \
  --memory 16384 \
  --disk path=cloud-config-${MY_NAME}.iso,device=cdrom \
  --disk path=ubuntu-18-${MY_NAME}.qcow2,device=disk --import \
  --network network=default,model=virtio \
  --graphics none \
  --noautoconsole
```

参数解释：
```
--name ubuntu-18-${MY_NAME}        虚拟机名称

--vcpus 8        分配8个cpu核心

--memory 16384    分配16G内存

--disk path=cloud-config-${MY_NAME}.iso,device=cdrom        使用前面步骤创建的 `cloud-config-${MY_NAME}.iso` 文件作为光盘

  --disk path=ubuntu-18-${MY_NAME}.qcow2,device=disk --import        使用前面步骤创建的 `ubuntu-18-${MY_NAME}.qcow2` 文件作为磁盘

--network network=default,model=virtio        使用NAT网络，只分配内部ip，目前使用足够了
```


3.2. 确认虚拟机正常运行

```sh
virsh list
```

根据命令行输出确认虚拟机正在运行状态

3.3. 查看虚拟机的NAT内部IP（动态）

```sh
virsh domifaddr ubuntu-18-${MY_NAME}
```

本例中得到的动态ip是 `192.168.122.119`

3.4. 使用ssh登入虚拟机，默认用户名为 `ubuntu`，密码为前面设置的`12345`

```
ssh ubuntu@192.168.122.119
```

登入虚拟机后，修改root用户密码（注意这是在虚拟机中操作）：

```sh
sudo passwd
```

根据提示，可以设置root用户密码为123abc。

修改ssh配置，允许ssh使用root用户登陆（注意这是在虚拟机中操作）：

```sh
sudo sed -i 's|^#PermitRootLogin prohibit-password$|PermitRootLogin yes|' /etc/ssh/sshd_config 
```

重新加载ssh配置（注意这是在虚拟机中操作）

```sh
sudo systemctl reload sshd
```

完成后`exit`退出ssh，重新以root用户登陆（密码123abc）：

```
ssh root@192.168.122.119
```

虚拟机的内部IP是DHCP分配的动态ip，需要设置为静态，方便后续使用（注意这是在虚拟机中操作）：

```sh
eval "$(ip route get 8.8.8.8 | sed -En 's|^.* via ([^ ]+) dev ([^ ]+) src ([^ ]+) .*$|GATEWAY=\1\nINTERFACE=\2\nIP=\3|p')"
cat >/etc/netplan/50-cloud-init.yaml <<EOF
network:
  version: 2
  ethernets:
    ${INTERFACE}:
      dhcp4: no
      addresses:
        - ${IP}/24
      gateway4: ${GATEWAY}
      nameservers:
        addresses: [223.5.5.5, 223.6.6.6]
EOF
```

注意：上面的脚本不需要修改，直接复制执行就可以。

确认无误（注意这是在虚拟机中操作）：

```sh
cat /etc/netplan/50-cloud-init.yaml
```

重启虚拟机（注意这是在虚拟机中操作，不要重启物理机）：

```sh
shutdown -r now
```

## 4. 暴露虚拟机 SSH(22) 端口

本例中的虚拟机ip `192.168.122.119` 是NAT网络分配的内部ip，只能在主机（物理机）上访问。

为了能从外部网络访问虚拟机，可以用socat命令把虚拟机的22端口转发到主机（物理机）上的端口:

```sh
docker run -d --name=ssh-${MY_NAME} --restart=always --net=host alpine/socat \
    tcp-listen:22119,fork,reuseaddr tcp4-connect:192.168.122.119:22
```

上述命令把 `192.168.122.119:22` 转发到主机上的 `22119` 端口，举例主机的ip为 `172.0.10.210` 可以在外部使用ssh命令登入虚拟机：

```sh
ssh -p 22119 root@172.0.10.210
```

## 5. 删除虚拟机

暴力关闭虚拟机

```sh
virsh destroy ubuntu-18-${MY_NAME}
```

删除虚拟机配置

```sh
virsh undefine ubuntu-18-${MY_NAME}
```

删除虚拟机镜像文件

```sh
rm ubuntu-18-${MY_NAME}.qcow2 cloud-config-${MY_NAME}.iso
```

删除用于端口转发的docker镜像

```sh
docker rm -f ssh-${MY_NAME}
```

================
docker run -it --rm \
  -v "${PWD}:/mnt/sdc1/weiguohong" \
  -w /mnt/sdc1/weiguohong \
  -u "$(id -u):$(id -g)" \
  --entrypoint /usr/bin/env \
  bkahlert/libguestfs \
  -- qemu-img resize ubuntu-18-wgh.qcow2 100G
  
  
dd if=/dev/zero bs=1G seek=100 count=0 of=ubuntu-18-wgh.qcow2

ls -alh ubuntu-18-wgh.qcow2


cat >cloud-config-wgh.txt <<EOF
#cloud-config
password: 12345
chpasswd: { expire: False }
ssh_pwauth: True
hostname: localhost
EOF

cloud-localds cloud-config-wgh.iso cloud-config-wgh.txt

  
virt-install \
  --name ubuntu-18-wgh \
  --os-type=linux \
  --vcpus 8 \
  --memory 16384 \
  --disk path=cloud-config-wgh.iso,device=cdrom \
  --disk path=ubuntu-18-wgh.qcow2,device=disk --import \
  --network network=default,model=virtio \
  --graphics none \
  --noautoconsole

NOTE:可以设置 network 为桥接模式:https://www.jianshu.com/p/51ce8858c35c
  
  
sudo sed -i 's|^#PermitRootLogin prohibit-password$|PermitRootLogin yes|' /etc/ssh/sshd_config

eval "$(ip route get 8.8.8.8 | sed -En 's|^.* via ([^ ]+) dev ([^ ]+) src ([^ ]+) .*$|GATEWAY=\1\nINTERFACE=\2\nIP=\3|p')"


cat >/etc/netplan/50-cloud-init.yaml <<EOF
network:
  version: 2
  ethernets:
    ens2:
      dhcp4: no
      addresses:
        - 192.168.122.168/24
        - 192:168:122:168:8e5:7dff:feac:e101/64
      gateway4: 192.168.122.1
      nameservers:
        addresses: [223.5.5.5, 223.6.6.6]
    ens5:
      dhcp4: false
      addresses:
        - 172.0.14.169/24
      gateway4: 172.0.14.1
      nameservers:
        addresses: [8.8.8.8, 114.114.114.114] 
EOF


network:
    ethernets:
        enp1s0:
            dhcp4: false
            addresses:
            - 172.0.14.53/24
            gateway4: 172.0.14.1
            nameservers:
                addresses:
                - 8.8.8.8
                - 114.114.114.114
                search: []
    version: 2




cat >/etc/netplan/50-cloud-init.yaml <<EOF
network:
    ethernets:
        ens2:
            dhcp4: false
            addresses:
            - 172.0.14.169/24
            gateway4: 172.0.14.1
            nameservers:
                addresses:
                - 8.8.8.8
                - 114.114.114.114
                search: []
    version: 2
EOF

network:
  ethernets:
    ens3:
        addresses:
        - 172.0.14.48/24
        dhcp4: false
        gateway4: 172.0.14.1
        nameservers:
            addresses:
            - 114.114.114.114
            search: []
    ens9:
        addresses:
        - 172.1.14.48/16
        dhcp4: false
  version: 2

# This file is generated from information provided by the datasource.  Changes
# to it will not persist across an instance reboot.  To disable cloud-init's
# network configuration capabilities, write a file
# /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg with the following:
# network: {config: disabled}
network:
  ethernets:
    ens2:
        addresses:
        - 172.0.14.48/24
        dhcp4: false
        gateway4: 172.0.14.1
        nameservers:
            addresses:
            - 114.114.114.114
            search: []
    ens9:
        addresses:
        - 172.1.14.48/16
        dhcp4: false
  version: 2
  
 ========给vm 添加网卡
 1) 物理主机上给vm 添加网卡
 https://blog.csdn.net/weixin_39094034/article/details/105373163
 2) 到vm 上给网卡添加ip //https://cloud-atlas.readthedocs.io/zh_CN/latest/linux/ubuntu_linux/network/netplan.html#:~:text=Netplan%E5%85%81%E8%AE%B8%E9%80%9A%E8%BF%87YAML%E6%8A%BD%E8%B1%A1%E6%9D%A5%E9%85%8D%E7%BD%AE%E7%BD%91%E7%BB%9C%E6%8E%A5%E5%8F%A3%EF%BC%8C%E5%9C%A8%20NetworkManager%20%E5%92%8C,systemd-networkd%20%E7%BD%91%E7%BB%9C%E6%9C%8D%E5%8A%A1%EF%BC%88%E5%BC%95%E7%94%A8%E4%B8%BA%20renderers%29%E7%BB%93%E5%90%88%E5%85%B1%E5%90%8C%E5%B7%A5%E4%BD%9C%E3%80%82
 
 e.g: 给ens5 添加ip
 cat >/etc/netplan/50-cloud-init.yaml <<EOF
network:
  version: 2
  ethernets:
    ens2:
      dhcp4: no
      addresses:
        - 192.168.122.168/24
        - 192:168:122:168:8e5:7dff:feac:e101/64
      gateway4: 192.168.122.1
      nameservers:
        addresses: [223.5.5.5, 223.6.6.6]
    ens5:
      dhcp4: false
      addresses:
        - 172.0.14.169/24
      gateway4: 172.0.14.1
      nameservers:
        addresses: [8.8.8.8, 114.114.114.114] 
EOF


3) netplant apply

4) may be to stop NAT network dev:
ifconfig [NIC_NAME] Down/Up
https://zhuanlan.zhihu.com/p/65480107
NOTE: vm 重起后，还是重新开启该网卡

5) 宿主机 移除vm 的网卡： https://www.zabbx.cn/archives/virsh%E6%96%B0%E5%A2%9E%E5%88%A0%E9%99%A4%E8%99%9A%E6%8B%9F%E6%9C%BA%E7%BD%91%E5%8D%A1

======================新vm 基础环境=============
### 1.win & vm 共享文件
```sh
1) install smab

2) updat smab config

3) service samb start

4) win connect vm

refer to: https://blog.csdn.net/hh3167253066/article/details/120528201
``` 
### 2. install git
```sh
1) apt install git

2) git connnect to github/gitlab: https://juejin.cn/post/6844904005152276494

```

### 3.install rke2 
```sh
refer to : ubuntu 虚拟机部署 rke2 环境  http://172.0.14.219:85/zh/Linux/deploy-rke2
```
NOTE: 
问题1：pull images  出现失败显示证书失败，linux 需要安装证书：
```sh
cat >/usr/share/ca-certificates/extra/rootCA.crt <<EOF
-----BEGIN CERTIFICATE-----
MIIDxjCCAq6gAwIBAgIQMf6njx2+HKhIOAtihJNAkzANBgkqhkiG9w0BAQsFADBV
MRMwEQYKCZImiZPyLGQBGRYDY29tMRwwGgYKCZImiZPyLGQBGRYMY2FzYS1zeXN0
ZW1zMSAwHgYDVQQDExdjYXNhLXN5c3RlbXMtQ0FTQUNBMS1DQTAeFw0yMjA5Mjkx
OTUyNDBaFw0yNzA5MjkyMDAyNDBaMFUxEzARBgoJkiaJk/IsZAEZFgNjb20xHDAa
BgoJkiaJk/IsZAEZFgxjYXNhLXN5c3RlbXMxIDAeBgNVBAMTF2Nhc2Etc3lzdGVt
cy1DQVNBQ0ExLUNBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA+TMv
7+HcU1od+eZudxrYWGX3I5wbOngAm0dTRyF51YfggiTtsixR9Ri6zwc4HMqETEV9
3x4RfziD6jE+4P3HNnf39TFZ5JPRy3NK3JbiVSBjJrlhhxQC3YfIdPqFIUWpcgxn
H1CFIh7AxDOqnBVITrAhZ79wElvwfuR2v9Z1lLL0QCSmuVc0hhUVgE/45T2LJUkX
sBfZDcWX47dj/UI18DsmnJqzeepeg5izwkuppuVUxAUBw9xsNNRScjaY2gv469eF
J4sBeQ5b2lF1nlZP+B0ju8o9/XxRp//Y0MmMG/hz6S7teXOgTifUvgG4GefYVezb
puewNHrSyaOgIiuDmwIDAQABo4GRMIGOMBMGCSsGAQQBgjcUAgQGHgQAQwBBMA4G
A1UdDwEB/wQEAwIBhjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBT+ezcNk5Vg
ph+Do6mpIx+pKb89zzASBgkrBgEEAYI3FQEEBQIDAgACMCMGCSsGAQQBgjcVAgQW
BBTbSllRujduMuzDI4Sv44PS+uyFsTANBgkqhkiG9w0BAQsFAAOCAQEAx+pcRuQ2
nv2+Quf+yxkVLs/0C7HdQb98gRMcnACSaktyNCt1mdodazrBHAlJZotn5aQUf+zp
nhlTh4FnNTfP4HKDoxdGUjGJFU3U3NH6uYl0ijeuoXeZ9YCFlnNslOUUsSSmJ/Ch
J8B8IYR3um6Do7MARn3Z56/dRPKMNcF6s8o4pGIHBlzRoxK98QunsWtp62tNaaGV
zcLSwdjKX3tivi3hTaxPRFeNjJ7HoUnJGpTPSZKmT+7VKSiHiplTJNAOPX4RFI6w
Y+56SJjNFbzZ1sr2JzB3RdVSsmWmpZOpEm76NB4kZGWP8I3QRNChaCpVbKuRolOx
55VVpbdAKoL0rw==
-----END CERTIFICATE-----
EOF

cat >>/etc/ca-certificates.conf<<EOF
/extra/rootCA.crt
EOF

sudo update-ca-certificates
```

### 4.搭建测试集群环境
现在跑 k8s 有三种方式：

1. robert 方式，部署 TMC 公共集群，跑 pipeline;
2. 手动测试目录1，也需要部署 TMC 公共集群，导出并修改成k8s可以直接apply的配置，然后跑测试；
3. 手动测试目录2，既可以使用 TMC 公共集群，也可以使用私人集群，跟之前 tekton 运行方式类似，唯一要做的就是把tmc的yaml配置转换成kubectl可以直接apply的配置（目前钉钉就是这种做法）
NOTE: 为何 统一使用方式1 ? 需要调研下可行性（之前没有用TEKTON是因为1. 耗时很久。 2. 很多问题是TEKTON自己引起的，多了太多维护工作



  