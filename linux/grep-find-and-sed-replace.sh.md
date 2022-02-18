##1 find and replace
```
find: 10000
replace: old:10000,  new:100
e.g:  grep "10000" * -R |awk -F: '{print $1}' |sort |uniq |xargs sed -i 's/10000/100/g'
```
```
方法一：find -name '要查找的文件名' | xargs perl -pi -e 's|被替换的字符串|替换后的字符串|g'
代码示例：

find -name 'pom.xml' | xargs perl -pi -e 's|http://repo1.maven.org/|http://registry.taobao.com/groups/public|g'

方法二：sed -i "s/原字符串/新字符串/g" `grep 原字符串 -rl 所在目录`
代码示例：

sed -i "s#10.220.96.205:8022#11.1.14.145#g" 'grep mahuinan -rl ./'

注：命令中的#可以替换成/或者|，以便于和字符串区分。

方法三：grep "原字符串" * -R | awk -F: '{print $1}' | sort | uniq | xargs sed -i 's/原字符串/新字符串/g'
代码示例：

grep "master" * -R | awk -F: '{print $1}' | sort | uniq | xargs sed -i 's/master/release/g'

批量替换配置文件中的IP：

grep "[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}" * -R | awk -F: '{print $1}' | sort | uniq | xargs sed -i 's/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/172\.0\.0\.1/g'

方法四： find 所在目录 -type f -path "文件名称"|xargs sed -i 's: 原字符串 : 新字符串 :g'
代码示例：

find ./ -name "*"|xargs grep "/data/" #查询匹配结果

find ./ -type f -path "*.sh"|xargs sed -i 's:/data/:/databak/:g'   #查找并替换

方法四：
sed -E -z -i 's/\n\s*x-kubernetes-list-map-keys:\n\s*- port\n\s*- protocol\n\s*x-kubernetes-list-type: map//g' operators/axyom/config/crd/bases/*.yaml

```

















