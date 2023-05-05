
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
        - 172.0.14.169/24  # 改为申请到的ip
        dhcp4: false
        gateway4: 172.0.14.1
        nameservers:
            addresses:
            - 114.114.114.114
            search: []
```


================
1、这里我们还是克隆kvm_client00，我们通过如下命令创建新虚拟机的配置文件

[root@5201351_kvm ~]# virsh dumpxml kvm_client00 > /etc/libvirt/qemu/kvm_client02.xml    //创建新虚拟机的配置文件
2、复制原虚拟机的磁盘文件，通过方法一、我们知道，磁盘默认位置为/var/lib/libvirt/images，我们执行如下命令进行复制

[root@5201351_kvm ~]# cd /var/lib/libvirt/images
[root@5201351_kvm images]# cp kvm_client00.img kvm_client02.img
3、直接编辑修改配置文件kvm_client02.xml，修改name,uuid,disk文件位置,mac地址，vnc端口
-uuid: https://uutool.cn/uuid/
-mac： https://blog.csdn.net/Mosicol/article/details/88294281

NOTE: refer: https://www.jianshu.com/p/cb1b6159966c
NOTE： 可能需要修改 https://www.cnblogs.com/liufarui/p/13144343.html

4、通过新虚拟机的配置文件，定义新的虚拟机，只需要执行如下一条命令即可。

[root@5201351_kvm ~]# virsh define /etc/libvirt/qemu/kvm_client02.xml   //通过配置文件定义新的kvm虚拟机
需要特别说明的是、以上两种方法克隆的虚拟机、我们都需要进入克隆的新虚拟机里

修改网卡设备文件/etc/udev/rules.d/70-persistent-net.rules，或者直接将其删除，再重启克隆的目的虚拟机

