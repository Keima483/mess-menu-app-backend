package models

type MenuItemModel struct {
	Day       string `json:"day"`
	BreakFast string `json:"breakfast"`
	Lunch     string `json:"lunch"`
	Snacks    string `json:"snacks"`
	Dinner    string `json:"dinner"`
}
