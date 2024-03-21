package middle

import "sort"

type SortRes struct {
	Docx  string
	Score float64
	Id    int
}

// 具体排序：
// qy为输入的query分词后的token形式，res则是搜索结构，返回值是将res排序好的结果。
func SortRess(qy []string, res []string) []*SortRes {
	exist := make(map[int]*SortRes)
	for _, v := range qy { // 遍历每一个query的分词后的token词条
		for i, v2 := range res { // 遍历每一个结果
			score := CalculateTFIDF(v, v2, res)
			// 记录分数构成，计算每个词条对每个文档结构的score
			if _, ok := exist[i]; !ok {
				// 如果exist中还没存在这个词条，则进行进行初始化
				tmp := &SortRes{
					Docx:  v2,
					Score: score,
					Id:    i,
				}
				exist[i] = tmp
			} else {
				// 如果已经存在了，则进行分数的相加
				// 意思就是每个res中的doc对于每个token的权重之和的结果。权重的对象始终都是res中doc
				exist[i].Score += score
			}
		}
	}
	resList := make([]*SortRes, 0)
	for _, v := range exist { // 构建结构体
		resList = append(resList, &SortRes{
			Docx:  v.Docx,
			Score: v.Score,
			Id:    v.Id,
		})
	}
	sort.Slice(resList, func(i, j int) bool { // 按照score进行排序
		return resList[i].Score > resList[j].Score
	})
	return resList
}

func Search(index InvertedIndex, query string, docs []string) ([]string, []string) {
	result := make(map[int]bool)
	qy := Tokenize(query)     // query词条进行分词
	for _, word := range qy { // 遍历分完词的每一个term
		if doc, ok := index[word]; ok {
			// 搜索倒排索引中，term对应的doc数组，doc数组就是存在该term词条的所有的doc id
			for _, d := range doc {
				// 对doc数组进行遍历，获取所有的doc id，并且进行标志。
				result[d] = true
			}
		}
	}

	output := []string{}
	for d := range result {
		output = append(output, docs[d])
		// 利用正排索引，找到id对应的存储内容并返回
	}
	return output, qy
}
