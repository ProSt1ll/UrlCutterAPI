package urlcut

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlCut_Shuffle(t *testing.T) {
	cut := New()
	var temp string
	var temp1 string
	cut1 := New()
	temp = string(cut.list)
	temp1 = string(cut1.list)
	fmt.Println(temp)

	assert.NotEqual(t, temp, temp1)
	for i := 0; i < 10; i++ {
		cut.Shuffle()
		cut1.Shuffle()

		assert.NotEqual(t, string(cut.list), temp)
		assert.NotEqual(t, string(cut1.list), temp1)
		assert.NotEqual(t, cut.list, cut1.list)
		temp = string(cut.list)
		temp1 = string(cut1.list)
	}
}

func TestUrlCut_CreateShortURL(t *testing.T) {
	cut := New()
	s := cut.CreateShortURL(0)
	temp := []rune("0000000000")
	for i, _ := range s {
		temp[i] = cut.list[0]
	}

	assert.Equal(t, s, string(temp))
	temp[len(temp)-1] = cut.list[1]
	assert.Equal(t, cut.CreateShortURL(1), string(temp))
}
