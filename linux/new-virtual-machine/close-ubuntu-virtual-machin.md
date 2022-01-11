
### 1. virt-tool clone vm
```
1. clone vm-image // 物理机上操作
2. create new vm base on vm-image // 物理机上操作
3. rest vm ip //进入创建新的 vm 设置 ip 
4. set vm user&passwd // 进入刚创建新的vm 操作
```
#### 2 准备克隆的镜像
```
vm 默认路径/var/lib/libvirt/images
```
```
virsh list // 获取物理主机上有哪些 vm
virsh edit xxx(vm name) // 查看vm 信息中 镜像存放的位置
```
#### 3 开始克隆镜像
```
$ mkdir /mnt/sdc1/new-vm-dir/
$ rsync --progress ubuntu18.04.img /mnt/sdc1/new-vm-dir/new-ubuntu18.04.img
```
#### 4 使用virt-manager创建新的虚拟机就可以了

#### 5 使用virt-manager 进入vm, 修改虚拟机的静态IP地址
```
1. vim打开/etc/netplan/01-netcfg.yaml
2. IP改为刚申请到空闲IP，保存后执行netplan apply或者重启虚拟机即可。
```
- 例如修改如下
```linux
# This file describes the network interfaces available on your system
# For more information, see netplan(5).
network:
  version: 2
  ethernets:
    enp1s0: # 有些机器clone过来后是ens3，保持原来的即可
        addresses:
        - 172.0.14.17/24  # 改为申请到的ip
        dhcp4: false
        gateway4: 172.0.14.1
        nameservers:
            addresses:
            - 114.114.114.114
            search: []
```

