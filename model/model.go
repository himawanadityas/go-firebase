package model

type MyData struct{
	Name string `json:"name"`	//`firestore:"name, omitempty"`
	Address string `json:"address"`     // `firestore:"address, omitempty"`
	Job string `json:"job"`//`firestore:"job, omitempty"`
}