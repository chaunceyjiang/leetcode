package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	if len(text) < 7 {
		return 0
	}
	balloonMap := map[string]int{"b": 0, "a": 0, "l": 0, "o": 0, "n": 0}

	for i := range text {
		_, ok := balloonMap[text[i:i+1]]
		if !ok {
			continue
		}
		balloonMap[text[i:i+1]]++
	}
	balloonMap["l"] = balloonMap["l"]/2
	balloonMap["o"] = balloonMap["o"]/2
	var count = 0

	for _, v := range balloonMap {
		if count == 0 {
			count = v
		} else if v < count {
			count = v

		}
	}

	return count
}

func main() {
	//fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
	fmt.Println(maxNumberOfBalloons("krhizmmgmcrecekgyljqkldocicziihtgpqwbticmvuyznragqoyrukzopfmjhjjxemsxmrsxuqmnkrzhgvtgdgtykhcglurvppvcwhrhrjoislonvvglhdciilduvuiebmffaagxerjeewmtcwmhmtwlxtvlbocczlrppmpjbpnifqtlninyzjtmazxdbzwxthpvrfulvrspycqcghuopjirzoeuqhetnbrcdakilzmklxwudxxhwilasbjjhhfgghogqoofsufysmcqeilaivtmfziumjloewbkjvaahsaaggteppqyuoylgpbdwqubaalfwcqrjeycjbbpifjbpigjdnnswocusuprydgrtxuaojeriigwumlovafxnpibjopjfqzrwemoinmptxddgcszmfprdrichjeqcvikynzigleaajcysusqasqadjemgnyvmzmbcfrttrzonwafrnedglhpudovigwvpimttiketopkvqw"))
}