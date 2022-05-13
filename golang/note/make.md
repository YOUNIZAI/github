chan需要调用make()创建后才可以使用，如果只是使用var ch chan int声明而不make，会产生deadlock：
https://segmentfault.com/a/1190000040024268
