# 1. io包是输入输出的核心抽象，它和os.File，bufio，net/http等很多地方结合使用
Go的“流式编程”就是以io.Reader /io.Writer接口为核心

# 2. io包的核心内容

## 2.1 核心接口
* io.Reader
* io.Writer
* io.Seeker
* io.Closer

## 2.2 常用函数
* io.Copy
* io.ReadFull
* io.ReadAll
* io.WriteString

## 3. 小练习
* 写一个函数，把一个文件的内容拷贝到另一个文件（用io.copy）
* 实现一个函数，读取标准输入（键盘输入），然后写入到文件
* 用io.WriteString给一个日志文件追加一行数据