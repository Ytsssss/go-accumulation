package algorithm

import "strconv"

type Codec struct {
	CodeMap map[string]string
	Count   int
}

func Constructor1() Codec {
	return Codec{
		CodeMap: make(map[string]string),
		Count:   1000,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	shortUrl := "http://tinyurl.com/"
	count := this.Count + 1
	shortUrl += strconv.Itoa(count)
	this.Count++
	if _, ok := this.CodeMap[shortUrl]; ok {
		this.encode(longUrl)
	}
	this.CodeMap[shortUrl] = longUrl
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.CodeMap[shortUrl]
}
