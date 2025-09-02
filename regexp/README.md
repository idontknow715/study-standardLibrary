# 1.regexp包是Go提供的正则表达式库，用来进行字符串的匹配、查找、替换。

# 2. regexp包核心内容
## 2.1 正则的编译
* regexp.Compile(pattern string): 编译正则，返回Regexp对象
* regexp.MustCompile(pattern string): 类似Compile，但如果正则非法会直接panic（常用于全局变量 ）

## 2.2 匹配
* MatchString(s string): 判断字符串是否匹配
* FindString(s string): 返回第一个匹配的字符串
* FindAllString(s string,n int): 返回所有匹配的字符串（n=-1,表示不限制个数）

## 2.3 分组匹配
* FindStringSubmatch(s string): 返回整个匹配+各分组匹配结果
* FindAllStringSubmatch(s string,n int): 返回所有分组匹配结果

## 2.4 替换
* ReplaceAllString(src,repl string):用字符串替换
* ReplaceAllStringFunc(src string, repl func(string)string): 用函数动态替换

## 2.5 分割
* Split(s string, n int ): 按照正则分割字符串

## 2.6 正则表达式的符号含义：
### 2.6.1 基本的元字符：
* `.`:匹配任意一个字符（除了换行）
* `\d`: 数字，相当于`[0-9]`
* `\w`: 字母、数字、下划线，相当于`[A-Za-z0-9]`
* `\s`: 空白字符（空格、tab、换行）
* `^`: 匹配开头
* `$`: 匹配结尾

### 2.6.2 数量词
* `*`: 重复0次或多次
* `+`: 重复1次或多次
* `?`: 重复0次或1次
* `{n}`:重复n次
* `{n,}`: 至少n次
* `{n,m}`: n到m次

### 2.6.3 分组和选择
* `(...)`: 分组
* `|`: 或
* 例子：
  * (dog|cat)匹配“dog”或“cat”
  * (ab)+匹配“ab”重复多次

### 2.6.4 字符集
* `[abc]`: 匹配a、b或c
* `[^abc]`: 除了a、b、c之外
* `[0-9]`:数字范围
* `[A-Za-z]`:大小写字母

### 2.6.5 逐步组合
* 邮箱：`^\w+@\w+\.\w+$`
* 手机号（中国）: `^1[3-9]\d{9}$`


# 3 小练习
* 邮箱校验：写一个正则，判断字符串是否是合法邮箱
* 提取数字：从“Order12345Number67890”中提取所有数字
* 替换敏感词：把“你是大傻瓜”里的“傻瓜”替换为“**”
* 分割字符串：用正则把“a,b;c|d”按照，；｜这些符号分割成[a,b,c,d]
