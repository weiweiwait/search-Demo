package middle

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 定义一个map结构，key 是一个 term，value 是包含有 term 关键字的文档的 id 数组。
type InvertedIndex map[string][]int

func FileOpen() []string {
	file, err := os.Open("data/movies.csv")
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()
	// 创建一个 Scanner 用来读取文件内容
	docx := make([]string, 0)
	scanner := bufio.NewScanner(file)
	// 逐行读取文件内容并打印
	for scanner.Scan() {
		re := make([]string, 0)
		line := scanner.Text()
		re = strings.Split(line, ",")
		if len(re) > 16 {
			docx = append(docx, re[16])
		}
	}
	if len(docx) > 0 {
		docx = docx[1:]
	}

	return docx
}

// 构建索引
func BuildIndex(docx []string) InvertedIndex {
	index := make(InvertedIndex)
	for i, d := range docx { // 遍历所有的docx
		for _, word := range Tokenize(d) { // 对所有的docx进行token
			if _, ok := index[word]; !ok { // 如果index不存在这个term了
				index[word] = []int{i} // 初始化并放入 行数
			} else {
				index[word] = append(index[word], i)
				// 如果index不存在，则放入该term所在的 行数，也就是 行数
			}
		}
	}

	return index
}
