package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
	return []byte(stamp), nil
}

func GetTinyUrl(url string) (urlData URL) {
	hashedUrl := md5.Sum([]byte(url))
	// taking encoded value to be stored as key
	encodedURL := base64.URLEncoding.EncodeToString(hashedUrl[:])

	urlData = URL{
		EncodedURL: encodedURL,
		LongURL:    url,
	}

	return
}
