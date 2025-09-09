# 1.encoding/json 提供了JSON和GO数据结构的相互转换,Go的struct可以用标签（tag）控制JSON字段名和序列化行为

# 2. 核心函数：
* 序列化(Go -> JSON)
  * json.Marshal(v) -> []byte
  * json.MarshalIndent(v,prefix,indent) -> 美化输出
* 反序列化(JSON -> Go)
  * json.Unmarshal(data,&v)

# 3.解码器与编码器
除了Marshal/Unmarshal，encoding/json还提供了流失处理：
* json.NewEncoder(io.Writer).Encode(v) -> 直接写到文件/网络中
* json.NewDecoder(io.Reader).Decode(&v) -> 直接从文件/网络读

# 小练习：
1. 基本序列化与反序列化
* 定义一个User结构体，把它转成JSON，再解析回来。

2. 使用标签控制JSON
* 给结构体加json:"字段名"、omitempty、- 试试看输出有什么不同

3. 从文件读取json配置
* 写一个config.json文件，内容是应用配置，用json.Decoder读出来

4.写JSON文件
* 用json.NewEncoder把日志写入log.json中
