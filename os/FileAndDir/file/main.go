package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	//f, err := os.Create("example.txt")
	//if err != nil {
	//	fmt.Println("创建文件失败,err:", err)
	//}
	//defer f.Close()
	//fmt.Println("创建文件成功")

	// 打开文件
	//f, err := os.Open("example.txt")
	//if err != nil {
	//	panic(err)
	//	fmt.Println("打开文件失败：")
	//}
	//fmt.Println("打开文件成功")
	//f.Close()
	//
	// 重命名文件
	//err := os.Rename("example.txt", "rename.txt")
	//if err != nil {
	//	panic(err)
	//	fmt.Println("重命名文件失败,err: ", err)
	//}
	//fmt.Println("重命名文件成功")

	//删除文件
	//err := os.Remove("rename.txt")
	//if err != nil {
	//	panic(err)
	//	fmt.Println("删除文件失败,err: ", err)
	//}
	//fmt.Println("删除文件成功")

	// 创建目录
	//err := os.Mkdir("mydir", 0755)
	//if err != nil {
	//	fmt.Println("创建目录失败")
	//} else {
	//	fmt.Println("创建目录成功")
	//}

	//err := os.MkdirAll("a/b/c", 0755)
	//if err != nil {
	//	fmt.Println("创建多级目录失败")
	//} else {
	//	fmt.Println("创建多级目录成功")
	//}

	//err := os.Remove("mydir")
	//if err != nil {
	//	fmt.Println("删除目录失败")
	//} else {
	//	fmt.Println("删除目录成功")
	//}

	err := os.RemoveAll("a")
	if err != nil {
		fmt.Println("递归删除失败")
	} else {
		fmt.Println("递归删除成功")
	}
}
