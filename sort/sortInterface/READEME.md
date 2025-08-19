如果要让自定义的类型支持通用排序，需要实现三个方法


{

    type Interface interface{

    Len()int
    Less(i,j int) bool   // 定义谁在前面（true表示i< j）
    Swap(i,j int) // 交换元素
    }
}