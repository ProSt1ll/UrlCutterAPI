package saver

import (
	"github.com/ProSt1ll/UrlCutterAPI/model"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestMemSaver_StoreURL(t *testing.T) {

	testURL, err := url.Parse("exampleLong")
	if err != nil {
	}
	s := NewMemSaver()
	testTable := []struct {
		item model.URLs
		want int
		name string
	}{
		{
			name: "OK1",
			item: model.URLs{
				LongUrl:  *testURL,
				ShortUrl: "exampleShort",
				Id:       1,
			},
			want: 0,
		},
		{name: "OK2",
			item: model.URLs{
				LongUrl:  *testURL,
				ShortUrl: "example2Short",
				Id:       1,
			},
			want: 0,
		},
		{name: "OK3",
			item: model.URLs{
				LongUrl:  *testURL,
				ShortUrl: "example3Short",
				Id:       1,
			},
			want: 0,
		},
	}

	for _, testCase := range testTable {
		result, err := s.StoreURL(testCase.item)
		assert.NoError(t, err)
		assert.Equal(t, result, testCase.want)

	}

}

func TestMemSaver_LoadLong(t *testing.T) {
	s1 := "http://example1Long.ru"
	s2 := "http://example2Long.ru"
	s3 := "http://example3Long.ru"
	testURL1, err := url.Parse(s1)
	if err != nil {

	}
	testURL2, err := url.Parse(s2)
	if err != nil {

	}
	testURL3, err := url.Parse(s3)
	if err != nil {

	}
	s := NewMemSaver()
	s.StoreURL(model.URLs{
		LongUrl:  *testURL1,
		ShortUrl: "exampleShort",
		Id:       2,
	})
	s.StoreURL(model.URLs{
		LongUrl:  *testURL2,
		ShortUrl: "example2Short",
		Id:       2,
	})
	s.StoreURL(model.URLs{
		LongUrl:  *testURL3,
		ShortUrl: "example3Short",
		Id:       2,
	})

	testTable := []struct {
		item string
		want model.URLs
		name string
	}{
		{
			name: "OK1",
			item: "exampleShort",
			want: model.URLs{
				LongUrl:  *testURL1,
				ShortUrl: "exampleShort",
			},
		},
		{
			name: "OK2",
			item: "example2Short",
			want: model.URLs{
				LongUrl:  *testURL2,
				ShortUrl: "example2Short",
			},
		},
		{
			name: "OK3",
			item: "example3Short",
			want: model.URLs{
				LongUrl:  *testURL3,
				ShortUrl: "example3Short",
			},
		},
	}

	for _, testCase := range testTable {
		result, ok := s.LoadLong(testCase.item)
		assert.True(t, ok)
		assert.Equal(t, result, testCase.want)

	}
}

func TestMemSaver_LoadShort(t *testing.T) {
	s1 := "http://example1Long.ru"
	s2 := "http://example2Long.ru"
	s3 := "http://example3Long.ru"
	testURL1, err := url.Parse(s1)
	if err != nil {

	}
	testURL2, err := url.Parse(s2)
	if err != nil {

	}
	testURL3, err := url.Parse(s3)
	if err != nil {

	}
	s := NewMemSaver()
	s.StoreURL(model.URLs{
		LongUrl:  *testURL1,
		ShortUrl: "exampleShort",
		Id:       2,
	})
	s.StoreURL(model.URLs{
		LongUrl:  *testURL2,
		ShortUrl: "example2Short",
		Id:       2,
	})
	s.StoreURL(model.URLs{
		LongUrl:  *testURL3,
		ShortUrl: "example3Short",
		Id:       2,
	})

	testTable := []struct {
		item url.URL
		want model.URLs
		name string
	}{
		{
			name: "OK1",
			item: *testURL1,
			want: model.URLs{
				LongUrl:  *testURL1,
				ShortUrl: "exampleShort",
			},
		},
		{
			name: "OK2",
			item: *testURL2,
			want: model.URLs{
				LongUrl:  *testURL2,
				ShortUrl: "example2Short",
			},
		},
		{
			name: "OK3",
			item: *testURL3,
			want: model.URLs{
				LongUrl:  *testURL3,
				ShortUrl: "example3Short",
			},
		},
	}

	for _, testCase := range testTable {
		result, ok := s.LoadShort(testCase.item)
		assert.True(t, ok)
		assert.Equal(t, result, testCase.want)

	}
}
