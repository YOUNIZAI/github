
1. Displays the directory size
```
du -d1 -h
```
2. tr: delete or replace the characters
```
e.g: All lower case letters are converted into upper case letters
cat testfile.txt |tr a-z A-Z 

```

3. convert 10 to 16
echo "obase=16;39" |bc


4. grep file

ls vzw_product_phase_1 |grep ^test_.*json$

5.添加特定字符到文件每行 行首 行尾
e.g:首行 添加 START, 行尾 添加 END
sed '/./{s/^/START&/;s/$/&END/}' file.txt 
link: https://www.cnblogs.com/aaronwxb/archive/2011/08/19/2145364.html
note: 's/$/&TAIL/g'中的字符g代表每行出现的字符全部替换，如果想在特定字符处添加，g就有用了，否则只会替换每行第一个，而不继续往后找了
