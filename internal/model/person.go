package model

type Person struct {
	Name string `json: name`

	SurName string `json: surname`

	Patronymic string `json: patronymic`
}
