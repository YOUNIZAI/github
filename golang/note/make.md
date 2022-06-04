Q: chan需要调用make()创建后才可以使用，如果只是使用var ch chan int声明而不make，会产生deadlock：
A:
chan 是复合类型(是一个结构体类型，需要用内置函数make 来初始化这个结构体里面成员； 反推: 知道它是结构体类型，如果只是简单声明，成员指针=nil, 后面直接使用 是为出问题的）
更多资料参考：https://segmentfault.com/a/1190000040024268
