package handlers

import (
	"encoding/json"
	model "go-firebase/model"
	config "go-firebase/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createData(w http.ResponseWriter, r *http.Request){
	var d model.MyData
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&d)
	if err != nil {
		log.Fatal(err)
	}
	
	ctx, app := config.Connection()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := client.Collection("profile").Doc("user").Set(ctx, d); err != nil {
		log.Fatal(err)
	}
	
	
	

}

func HandlerHttp(){
	// os.Setenv("PORT","8080")
	// port := os.Getenv("PORT") // for deploy
	router := mux.NewRouter()
	router.HandleFunc("/gofirebase/add", createData).Methods("POST")
	// router.HandleFunc("/mth/nestle/{id}", getSingleData).Methods("GET")
	// router.HandleFunc("/mth/nestle/status/{status}", getStatusData).Methods("GET")
	// router.HandleFunc("/mth/nestle/filename/{file_zip}", getFileName).Methods("GET")
	http.ListenAndServe(":8080", router)  // for local
	// http.ListenAndServe(":"+port, router) // for deploy

// 	srv := &http.Server{
// 		Handler: router,
// 		Addr: "mth-ftp-apps.herokuapp.com",
// 	}
// 	log.Fatal(srv.ListenAndServe())
}