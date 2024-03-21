package middle

import (
	"github.com/go-ego/gse"
	"strings"
)

var StopWord = []string{",", ".", "。", "*", "(", ")", "'", "\""}
var gobalGse gse.Segmenter

// 进行数据清洗
func RemoveShopWord(word string) string {
	for i := range StopWord {
		word = strings.Replace(word, StopWord[i], "", -1)
	}
	return word
}

// 分词
func InitConfig() {
	gobalGse, _ = gse.New()
}
func Tokenize(text string) []string {
	text = RemoveShopWord(text)
	return gobalGse.CutSearch(text)
}
