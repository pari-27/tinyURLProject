package utils

type Response struct {
	Status bool     `json:"status"`
	Error  ErrorObj `json:"error,omitempty"`
	Data   DataObj  `json:"data,omitempty"`
}
type DataObj struct {
	ShortUrl string `json:"shortUrl,omitempty"`
	Message  string `json:"message,omitempty"`
}
type ErrorObj struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
