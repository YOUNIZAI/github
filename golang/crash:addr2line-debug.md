### method 1
思路：
```
1. 获取可执行文件：方法1： 自己编译源代码， 脚本：build_docker.sh  方法2： 从镜像里面获取（这个可以查资料 了解一下）  
2. 通过log 找到大概 crash 堆栈地址    
3. 命令 addr2line -Cfe <可执行文件> <地址>
```
```
从镜像里面获取 步骤：
1.docker pull registry.gitlab.casa-systems.com/mobility/smf/sm:xxx
2.sudo docker run -it registry.gitlab.casa-systems.com/mobility/smf/sm:xxx sh
3. 新打开一个命令终端，执行：
 sudo docker cp 容器ID:/opt/casa/smf/sm/smfsm ./
 === note: 容器ID 通过docker ps 来找到
4. addr2line -Cfe ./smfsm 45fa43

```

Now we saw this in SQA's test env.
1. docker pull image registry.gitlab.casa-systems.com/mobility/smf/sm:xxx

2. Run the image if container is not running anymore.
<br>casa@mec7:~$ sudo docker run -it registry.gitlab.casa-systems.com/mobility/smf/sm:xxx sh

3. casa@mec7:~$ sudo docker ps
<br>CONTAINER ID        IMAGE                                                               COMMAND                   CREATED             STATUS              PORTS               <br>NAMES
<br>90a1d369437d        registry.gitlab.casa-systems.com/mobility/smf/sm:xxx                    "./smfsm"                 2 hours ago         Up 2 hours                              
4. Copy smfsm to VM which have addr2line tool.
<br>"docker cp <container id/name>:<the path of executable file>  <copy to where>"
<br>casa@mec7:~$ sudo docker cp 90a1d369437d:/opt/casa/smf/sm/smfsm ./
 
5. Using addr2line to find the line.
<br>addr2line -Cfe <the path of executable file>  <hex of where crash(backtrace)>
<br>casa@mec7:~$ addr2line -Cfe ./smfsm 45fa43          

### method 2
 Thank you for the crash debugging steps! In cases where we have a good crash stack we can also use disassemble -l which I had sent before for debugging crash.
 
Example: 000000c000f74928:  00000000010e2ad6 <gitlab.casa-systems.com/mobility/smf/sm/pkg/fsm.buildPfcpEstReqForUpfDefFlow+646>  0000000000010246
 
root@df8e4f759fe6:/opt/casa/smf/sm# ./dlv exec ./smfsm
Type 'help' for list of commands.
(dlv) disassemble -l buildPfcpEstReqForUpfDefFlow
TEXT gitlab.casa-systems.com/mobility/smf/sm/pkg/fsm.buildPfcpEstReqForUpfDefFlow(SB) /opt/casa/smf/sm/pkg/fsm/sess_upf.go
        sess_upf.go:1157                            0x10e28d0          64488b0c25a8ffffff                 mov rcx, qword ptr fs:[0xffffffa8]
        sess_upf.go:1157                            0x10e28d9          488d842480feffff                     lea rax, ptr [rsp+0xfffffe80]
           sess_upf.go:1157                            0x10e28e1          483b4110                                      cmp rax, qword ptr [rcx+0x10]
 
Looks for 10e2ad6 to find line number.
 
### method 3
Add addr2line tools into SMF image.
Dockerfile:
```
WORKDIR /opt/xxx
#install addre2line
COPY tools/addr2line addr2line
# Run apk addr2line 
COPY --from=build /opt/xxx/addr2line/addr2line /user/local/bin/
COPY --from=build /opt/xxx/addr2line/lib64 /lib64
COPY --from=build /opt/xxx/addr2line/x86_64-linux-gun  /lib/x86_64-linux-gun
```


Learning how to using golang addr2line tool.
Re: usage: echo ${stack_address} | go tool addr2line ${project_binary_flie}
Example:
000000c000f74928:  00000000010e2ad6 <gitlab.casa-systems.com/mobility/smf/sm/pkg/fsm.buildPfcpEstReqForUpfDefFlow+646>  0000000000010246
echo 00000000010e2ad6 | go tool addr2line smfsm
Output:
          gitlab.casa-systems.com/mobility/smf/sm/pkg/fsm.UpdateSdFromEstbRsp
          /root/git/src/gitlab.casa-systems.com/mobility/smf/sm/pkg/fsm/sess_upf.go:962
