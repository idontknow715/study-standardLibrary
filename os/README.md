# 1. os包是Go里非常核心的一个包，用来和操作系统交互。

# 2. os包的核心内容
## 2.1 文件和目录操作
* os.Open,os.Create,os.Remove,os.Rename
* os.Mkdir,os.MkdirAll,os.RemoveAll

## 2.2 文件读写
* File.Read,File.Write,File.Seek,File.Close

## 2.3 环境变量
* os.Getenv,os.Setenv,os.Environ

## 2.4 进程相关
* os.Exit,os.Getpid,os.Getppid

## 3 小练习
## 3.1 读取文件内容
* 写一个程序，打开一个文本文件，打印里面的所有内容

## 3.2 写入日志文件
* 写一个日志函数，把当前时间和一条信息写入到app.log文件中，每次运行都追加内容

## 3.3 环境变量练习
* 获取环境变量PATH，并把它打印出来
* 给自己设置一个环境变量MYNAME=xxx，并在程序中读取