package model

type Car struct {
	RegNum string `json: regNum`
	Mark   string `json: mark`
	Model  string `json: model`
	Owner  Person `json: owner`
	Year   int    `json: year`
}
