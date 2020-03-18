package main

import "fmt"

func countCharacters(words []string, chars string) int {
	// 一个hash存储字母表中每个字母出现次书
	cMap := make(map[rune]int)
	for _, c := range chars {
		cMap[c]++
	}
	// 初始化长度和
	res := 0
	for _, word := range words {
		// 遍历单词中的字母
		// 遍历字母时，如果有字母不在字母表中，字母不满足，直接跳出，否则计数
		// charMatch字母是否满足
		// strMatch单词是否满足
		charMatch := true
		strMatch := true
		thisMap := map[rune]int{}
		for _, w := range word {
			if cMap[w] > 0 {
				thisMap[w]++
			} else {
				charMatch = false
				break
			}
		}

		// 如果字母满足，遍历单词出现的字母，如果每个字母都满足，则单词满足，增加长度
		if charMatch {
			for c, count := range thisMap {
				if count > cMap[c] {
					strMatch = false
					break
				}
			}
			if strMatch {
				fmt.Printf("%s\n", word)
				res += len(word)
			}
		}
	}
	return res
}

// times: 1

func main() {
	//words := []string{"cat","bt","hat","tree"}
	//st := "atach"

	words := []string{
		"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnpznjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin", "ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb", "ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl",
		"boygirdlggnh",
		"xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx", "nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop", "hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx", "juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobiwzxnhdkidr", "lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo", "oxgaskztzroxuntiwlfyufddl", "tfspedteabxatkaypitjfkhkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp", "qnagrpfzlyrouolqquytwnwnsqnmuzphne", "eeilfdaookieawrrbvtnqfzcricvhpiv", "sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz", "yhouqhjevqxtecomahbwoptzlkyvjexhzcbccusbjjdgcfzlkoqwiwue", "hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpuvgiyeudhncv", "cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo", "teyygdmmyadppuopvqdodaczob", "qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs", "qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonrhcsubchs",
	}
	st := "usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp"
	fmt.Println(countCharacters(words, st))
}
