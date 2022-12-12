package urlcut

import (
	"math/rand"
	"time"
)

type UrlCut struct {
	list      []rune
	count     int
	generator *rand.Rand
}

func New() *UrlCut {
	temp := UrlCut{
		list:      []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrsuvwxyz_"),
		count:     len("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrsuvwxyz_"),
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	temp.Shuffle()
	return &temp
}

// Shuffle the list of symbols
func (u *UrlCut) Shuffle() {

	for i, _ := range u.list {
		random := u.generator.Intn(len("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrsuvwxyz_"))
		u.list[i], u.list[random] = u.list[random], u.list[i]
	}
}

// 1982608315404440064116146708361898137544773690227268628106279599612729753600000000000000 max cnt of uuid

// CreateShortURL create unique string from int
func (u *UrlCut) CreateShortURL(val int) string {
	var n int
	s := []rune("0000000000")
	for i, _ := range s {
		s[i] = u.list[0]
	}
	var cnt int = 0
	for val > 0 {
		n = val % u.count
		s[len(s)-cnt-1] = u.list[n]
		val = val / u.count
		cnt++
	}

	return string(s)

}
