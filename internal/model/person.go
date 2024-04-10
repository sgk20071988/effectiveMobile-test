package model

// Person description.
// swagger:model person
type Person struct {
	// Firstname of the person
	//
	// required: true
	Name string `json: name`
	// surname of the person
	//
	// required: true
	SurName string `json: surname`
	// Patronymic of the person
	//
	// required: false
	Patronymic string `json: patronymic`
}
