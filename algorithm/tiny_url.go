package algorithm

type Codec struct {
	d_map  string
	urlMap map[string]string
}

func Constructor() Codec {
	return Codec{d_map: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		urlMap: make(map[string]string)}

}

func (this *Codec) decimalTo62(hash int) string {
	var encoded string
	for ; hash != 0; {
		encoded = string(this.d_map[hash%len(this.d_map)])+encoded
		hash/=len(this.d_map)
	}
	return encoded
}
func hash(url string) int {
	hash := 0
	for i := 0; i < len(url); i++ {
		hash += int(url[i]) ^ 31
	}
	return hash
}

// Encodes a URL to a shortened URL.
func (this *Codec) Encode(longUrl string) string {
	_hash := hash(longUrl)
	var encoded string
	for {
		encoded = this.decimalTo62(_hash)
		if len(this.urlMap[encoded]) == 0 {
			this.urlMap[encoded] = longUrl
			break
		} else {
			_hash++
		}
	}
	return encoded
}

// Decodes a shortened URL to its original URL.
func (this *Codec) Decode(shortUrl string) string {
	if len(this.urlMap[shortUrl]) > 0 {
		return this.urlMap[shortUrl]
	}
	return ""
}
