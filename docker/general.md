
1. docker log 太多导致满盘
  links:https://colobu.com/2018/10/22/no-space-left-on-device-for-docker/

2. 删除 docker volume(volume 太多 会消耗多余 空间)
(link:) [https://johng.cn/docker-disk-usage-analyse-and-clean/#i] 
-  docker system df 
-  docker volume rm $(docker volume ls -qf dangling=true)

3. docker rmi more images
docker images  |grep smf/sm |grep gcs |awk   '{print }' | xargs docker rmi 

