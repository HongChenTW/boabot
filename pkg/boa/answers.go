package boa

import (
	"math/rand"
	"time"
)

const (
	defaultQuestion = "(藏在心裡的話)"
	defaultAnswer   = "順從你的心"
)

var (
	answers = []string{
		// original answers
		"為做最好的決定，保持冷靜",
		"別忘了生活還有樂趣",
		"轉移你的注意力",
		"援助將使你的發展取得成功",
		"不要指望它",
		"結果將取決於你的選擇",
		"把注意力集中在細節上",
		"試著改變你的日常生活",
		"你會因為你做了而感到快樂的",
		"絕不",
		"採用一種大膽的態度",
		"等待",
		"還有很多事等著你努力",
		"一年以後，已經不重要了",
		"跟隨別人的領導",
		"準備迎接意外",
		"是的",
		"清除障礙",
		"去遵循專家的建議",
		"懷疑它吧",
		"它可能是非凡的",
		"你不是真心的",
		"晚一點再處理",
		"負責",
		"實際一點",
		"隨心而行",
		"情況很快就會有改變",
		"你必須現在行動",
		"保持靈活",
		"是時候了",
		"晚一點再處理",
		"相信自己獨到的想法",
		"為什麼不做一個計畫",
		"看看會發生什麼",
		"可能會發生結果令人吃驚的事件",
		"要耐心",
		"別浪費你的時間",
		"這不是一件普通的事，慎重考慮",
		"絕對不行",
		"無論如何",
		"想想什麼才是重要的",
		"答案就在你的後院",
		"無關緊要，時間沖淡一切",
		"你將不得不妥協",
		"不要再問了",
		"它會影響到別人對你的看法",
		"它肯定會讓事情更有趣",
		"最好等一等",
		"現在是制定計劃的好時機",
		"你需要收集更多資訊",
		"必須的",
		"現在不要再提問題了，神有些累了",
		"如果你不抗拒，可以這麼幹",
		"最好等一等",
		"這不明智",
		"要謹慎",
		// extended answers
		"你冷靜",
		"在有跟沒有之間",
		"尋求援助吧",
		"不要指望 TA",
		"魔鬼藏在細節裡",
		"BJ4",
		"就做吧",
		"毋湯",
		"大膽一點吧",
		"你就等",
		"放眼世界，征服宇宙",
		"可以問你媽",
		"有 87% 可能性",
		"謀恩丟",
		"如果你能一個禮拜不喝飲料的話",
		"可以去問石頭",
		"結果可能還不錯",
		"了不起，負責",
		"你想幹嘛就幹嘛",
		"拖延一下沒關係",
		"可能會發生令人吃驚的事",
		"答案就在你身後",
		"這種小事可以不用問",
		"如果你不討厭，就做吧",
	}
)

func init() {
	rand.Seed(time.Now().Unix())
}

// GetAnswer returns a random answers
func GetAnswer() string {
	answersCount := len(answers)
	if answersCount == 0 {
		return defaultAnswer
	}
	return answers[rand.Intn(answersCount)]
}