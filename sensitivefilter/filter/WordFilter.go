package WordFilter

import (
	"fmt"
)

type wordtree struct {
	tree map[string]*wordtree
	isEnd bool
}

/*
* this is store sensitive words in memory
* private variable 
*/
var sensitiveWordMap wordtree

/*
* local sensitive word list to memory
*/
func loadSensitiveWord(set []string) {
	for _ , value := range sensitiveWords {
		nowMap := &sensitiveWordMap
		wordLenth := len(value)
		for i := 0 ; i < wordLenth ; i ++ {
			char := string(value[i])
			wordMap := nowMap.tree[char]
			if wordMap != nil {  
				nowMap = wordMap
			}else{
				var newtree wordtree
				if(nowMap.tree == nil){
					nowMap.tree = make(map[string]*wordtree)
				}
				nowMap.tree[char] = &newtree
				nowMap = &newtree
			}
			if(i == wordLenth-1){
				nowMap.isEnd = true
			}
		}
	}
}

func addSensitiveWord (){

}
func delSensitiveWord (){

}
func mofiSensitiveWord (){

}
