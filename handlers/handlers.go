package handlers

import (
	"encoding/json"

	config "go-firebase/config"
	model "go-firebase/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

func createData(w http.ResponseWriter, r *http.Request){
	var d model.MyData
	dataView := []model.MyData{}
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

	if _, err := client.Collection("profile").NewDoc().Set(ctx, d); err != nil {
		log.Fatal(err)
	}

	iterData := client.Collection("profile").Documents(ctx)
	for {
		data, err := iterData.Next()
		if err == iterator.Done {
			break
		}
		if err !=  nil {
			log.Fatal(err)
		}
		jsonBody, err := json.Marshal(data.Data())
		if err != nil {
			log.Fatal(err)
		}
		
		if err := json.Unmarshal(jsonBody, &d); err != nil {
			log.Fatal(err)
		}
		dataView = append(dataView, d)
	}
	json.NewEncoder(w).Encode(dataView)
	
}

func getAllData(w http.ResponseWriter, r *http.Request){
	var d model.MyData
	dataView := []model.MyData{}
	ctx, app := config.Connection()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	iterData := client.Collection("profile").Documents(ctx)
	for {
		data, err := iterData.Next()
		if err == iterator.Done {
			break
		}
		if err !=  nil {
			log.Fatal(err)
		}
		jsonBody, err := json.Marshal(data.Data())
		if err != nil {
			log.Fatal(err)
		}
		
		if err := json.Unmarshal(jsonBody, &d); err != nil {
			log.Fatal(err)
		}
		dataView = append(dataView, d)
	}
	json.NewEncoder(w).Encode(dataView)
}

func HandlerHttp(){
	// os.Setenv("PORT","8080")
	// port := os.Getenv("PORT") // for deploy
	router := mux.NewRouter()
	router.HandleFunc("/gofirebase/add", createData).Methods("POST")
	router.HandleFunc("/gofirebase/all", getAllData).Methods("GET")
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