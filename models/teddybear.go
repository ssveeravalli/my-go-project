package models

type TeddyBear struct {
	Name           string  `json:"name"`
	Color          string  `json:"color"`
	Occupation     string  `json:"occupation"`
	Characteristic string  `json:"characteristic"`
	Age            float32 `json:"age"`
}
