# Golang和Swoole,在http中的性能对比

## 系统环境

* MacBook Pro (Retina, 13-inch, Early 2015)
* 处理器 2.7 GHz Intel Core i5
* 内存 8 GB 1867 MHz DDR3
* go版本 go version go1.11.4 darwin/amd64
* PHP 7.1.1 (cli) (built: Jan 20 2017 07:19:55) ( NTS )
* Swoole 1.9.10

## 代码

https://github.com/kphcdr/golandvsswoole

## 测试方法

1. ab -n 2000 -c 2000  http://127.0.0.1:9001/
2. 等5分钟 
3. ab -n 2000 -c 2000  http://127.0.0.1:9002/
4. 重复十次

## 测试结果

![测试结果](图片链接 "")






























