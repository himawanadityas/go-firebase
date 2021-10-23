package model

type MyData struct{
	Name string	`firestore:"name, omitempty"`
	Address string `firestore:"address, omitempty"`
	Job string `firestore:"job, omitempty"`
}