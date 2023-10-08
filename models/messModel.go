package models

type MessModel struct {
	MessName      string `json:"messName"`
	BreakFastTime string `json:"breakFastTime"`
	LunchTime     string `json:"lunchTime"`
	SnacksTime    string `json:"snacksTime"`
	DinnerTime    string `json:"dinnerTime"`
}
