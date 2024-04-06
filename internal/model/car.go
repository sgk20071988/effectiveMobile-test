package model

type Car struct {
	RegNum string `json: regNum`
	Mark   string `json: mark`
	Model  string `json: model`
	Owner  string `json: owner`
}
