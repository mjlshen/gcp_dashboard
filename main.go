package main

import (
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/webview/webview"

	"github.com/mjlshen/mik8s/v2/pkg/gcloud"
)

// main function
func main() {
	folder := packr.New("React Front End", "./ui/build")
	http.Handle("/", http.FileServer(folder))
	go http.ListenAndServe(":8081", nil)

	go func() {
		r := Router()
		log.Fatal(http.ListenAndServe(":8082", r))
	}()

	// create a web view
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("mik8s")
	w.SetSize(1000, 800, webview.HintNone)
	w.Navigate("http://localhost:8081")
	w.Run()
}

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/gcloud/buckets", gcloud.GetBuckets).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/gcloud/projects", gcloud.GetProjects).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/gcloud/clusters", gcloud.GetGKEClusters).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/gcloud/serviceAccounts", gcloud.GetServiceAccounts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/gcloud/iam", gcloud.GetIamPolicyBindings).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	// router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
