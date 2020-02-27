package model

import "encoding/json"

type Article struct {
	Title string
	Content string
	DateTime string
	Source string
	Stock []Stock
	Plate []Plate
}
type Stock struct {
	Name string
	Code string
	Price string
	Float string
}

type Plate struct {
	Name string
	Float string
	// Trend string
	// Message string
	// MessageDate string
}

func FromJson(c interface{}) (Article,error) {

	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return Article{}, err
	}
	article := Article{}
	err = json.Unmarshal(jsonBytes, &article)
	if err !=nil {
		return Article{},err
	}
	return article,nil

}