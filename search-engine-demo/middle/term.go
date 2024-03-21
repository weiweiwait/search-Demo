package middle

import "strings"

// 计算TF
func CalculateTF(term string, document string) float64 {
	termCount := strings.Count(document, term)
	totalWords := len(Tokenize(document))
	return float64(termCount) / float64(totalWords)
}

// 计算IDF
func CalculateIDF(term string, documents []string) float64 {
	docWithTerm := 0
	for _, doc := range documents {
		if strings.Contains(doc, term) {
			docWithTerm++
		}
	}
	return float64(len(documents)) / float64(docWithTerm)
}

// TF*IDF即可获取权重，下面这里是由于数据问题，我是乘以100的
func CalculateTFIDF(term string, document string, documents []string) float64 {
	tf := CalculateTF(term, document)
	idf := CalculateIDF(term, documents)
	return tf * idf * 100.0
}
