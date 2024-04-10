package model

// Person description.
// swagger:model car
type Car struct {
	// Registration number of the car
	//
	// required: true
	RegNum string `json: regNum`
	// Mark of the car
	//
	// required: true
	Mark string `json: mark`
	// Model of the car
	//
	// required: true
	Model string `json: model`
	// Owner of the car
	//
	// required: true
	Owner Person `json: owner`
	// Year of manufacture of the car
	//
	// required: true
	Year int `json: year`
}
