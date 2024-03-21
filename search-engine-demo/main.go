package main

import (
	"fmt"
	"se/middle"
)

func main() {
	query := "王小波,徐克"
	middle.InitConfig() // 初始化配置
	docx := middle.FileOpen()
	index := middle.BuildIndex(docx) // 创建index
	res, qy := middle.Search(index, query, docx)
	fmt.Printf("一共%d记录，query分词结果%v\n", len(res), qy)
	resList := middle.SortRess(qy, res)
	for i := range resList {
		fmt.Println(resList[i].Score, resList[i].Docx)
	}
}
