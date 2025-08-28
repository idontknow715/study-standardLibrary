# 1. bufio包是什么
bufio包的核心作用就是给I/O加一层缓冲，这样能减少系统调用的次数，让读写更高效

比如：读文件时不是一字节一字节去读，而是一次性读一块缓冲区，写入时也类似。

# 2. bufio核心内容

## 2.1 bufio.Reader 
* 用于高效读取，可以一行一行、一个词一个词地读
* 常用方法：
    * ReadString(delim byte): 读到指定分隔符
    * ReadBytes(delim byte): 类似上面，不过返回的是[]byte
    * Peek(n int): 预读n个字节，不移动游标

## 2.2 bufio.Writer
* 用于高效写入，先写到缓冲区，再一次性写到目标
* 常用方法：
  * WriteString(s string): 写入字符串
  * Write(p []byte): 写入字节切片
  * Flush(): 把缓冲区内容真正写出去，一定要调用它

## 2.3 bufio.Scanner
* 用来一行一行地扫描输入，非常常用
* 常用方法：
  * Scan(): 移动到下一行（返回true/false）
  * Text(): 返回当前行内容

## 3. 小练习
* 逐行复制文件：写一个程序，用 bufio.Reader 读取 input.txt，逐行写入到 output.txt。
* 追加写入日志：写一个函数 Log(msg string)，把 msg 和当前时间写入 app.log 文件，每次调用都追加到文件末尾。
* 单词统计器：用 bufio.Scanner 从标准输入读取内容（多行），直到输入 exit。
  统计输入的单词总数，并打印结果。
* 过滤文本：读取一个文件，逐行扫描，如果一行中包含 "ERROR"，就写入 error.log 文件，否则丢弃。