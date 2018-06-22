package models

type Splash struct {
	Title     string `json:"title"`
	Pic       string `json:"pic"`
	PicX      string `json:"pic_x"`
	WebUrl    string `json:"web_url"`
	CountDown int64  `json:"count_down"`
}
