package gcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/iterator"
)

// func init() {
// 	go func() {
// 		r := Router()
// 		log.Fatal(http.ListenAndServe(":8080", r))
// 	}()
// }

// func Router() *mux.Router {

// 	router := mux.NewRouter()

// 	router.HandleFunc("/api/v1/gcloud/buckets", GetBuckets).Methods("GET", "OPTIONS")
// 	router.HandleFunc("/api/v1/gcloud/projects", GetProjects).Methods("GET", "OPTIONS")
// 	// router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
// 	// router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
// 	// router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
// 	// router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
// 	// router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
// 	return router
// }

func listProjects() *[]Project {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	cloudresourcemanagerService, err := cloudresourcemanager.New(c)
	if err != nil {
		log.Fatal(err)
	}

	var projects []Project
	req := cloudresourcemanagerService.Projects.List()
	if err := req.Pages(ctx, func(page *cloudresourcemanager.ListProjectsResponse) error {
		for _, project := range page.Projects {
			if project.LifecycleState == "ACTIVE" {
				p := new(Project)
				p.ProjectId = project.ProjectId
				p.Number = project.ProjectNumber
				projects = append(projects, *p)
			}
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return &projects
}

// ProjectGenerator takes in a variable number of projects and converts them into a channel
// following the "pipeline" architecture. This is an entrypoint function to the rest
// of the functions processing Endpoints and can be interrupted/will trigger downstream
// functions to stop processing with the done channel.
func ProjectGenerator(done <-chan interface{}, projects ...Project) <-chan Project {
	pc := make(chan Project)
	go func() {
		defer close(pc)
		for _, p := range projects {
			select {
			case <-done:
				return
			case pc <- p:
			}
		}
	}()
	return pc
}

// BucketGenerator takes in a variable number of buckets and converts them into a channel
// following the "pipeline" architecture. This is an entrypoint function to the rest
// of the functions processing Endpoints and can be interrupted/will trigger downstream
// functions to stop processing with the done channel.
func BucketGenerator(done <-chan interface{}, buckets ...Bucket) <-chan Bucket {
	bc := make(chan Bucket)
	go func() {
		defer close(bc)
		for _, b := range buckets {
			select {
			case <-done:
				return
			case bc <- b:
			}
		}
	}()
	return bc
}

func BucketAssembler(done <-chan interface{}, bc <-chan Bucket) []Bucket {
	var buckets []Bucket

	for b := range bc {
		buckets = append(buckets, b)
	}

	return buckets
}

func listBuckets(done <-chan interface{}, pc <-chan Project) []Bucket {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var buckets []Bucket

	for p := range pc {
		it := client.Buckets(ctx, p.ProjectId)
		for {
			battrs, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Printf("In project_id: %s - %v", p.ProjectId, err)
				break
			}
			b := new(Bucket)
			b.Id = battrs.Name
			b.Name = battrs.Name
			b.ProjectId = p.ProjectId
			b.StorageClass = battrs.StorageClass
			b.VersioningEnabled = battrs.VersioningEnabled
			b.Location = battrs.Location
			buckets = append(buckets, *b)
		}
	}

	return buckets
}

func listGKEClusters(done <-chan interface{}, pc <-chan Project) []GKECluster {
	// consoleBaseURL := "https://console.cloud.google.com/kubernetes/clusters/details"
	// clusterURL := fmt.Sprintf("%v/%v/%v/?project=%v\n", consoleBaseURL, cluster.Location, cluster.Name, s.Text)
	var clusters []GKECluster

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	containerService, err := container.New(c)
	if err != nil {
		log.Fatal(err)
	}

	for p := range pc {
		parent := fmt.Sprintf("projects/%s/locations/-", p.ProjectId)

		resp, err := containerService.Projects.Locations.Clusters.List(parent).Context(ctx).Do()
		if err != nil {
			log.Println(err)
			break
		}

		for _, cluster := range resp.Clusters {
			c := new(GKECluster)
			c.Id = fmt.Sprintf("%s-%s", p.ProjectId, cluster.Name)
			c.ProjectId = p.ProjectId
			c.Name = cluster.Name
			c.Location = cluster.Location
			c.CurrentMasterVersion = cluster.CurrentMasterVersion
			c.CurrentNodeVersion = cluster.CurrentNodeVersion
			clusters = append(clusters, *c)
		}
	}

	return clusters
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	projects := listProjects()
	json.NewEncoder(w).Encode(projects)
}

func GetGKEClusters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	projects := listProjects()
	done := make(chan interface{})
	defer close(done)
	pc := ProjectGenerator(done, *projects...)
	clusters := listGKEClusters(done, pc)

	json.NewEncoder(w).Encode(clusters)
}

func GetBuckets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	projects := listProjects()
	done := make(chan interface{})
	defer close(done)
	pc := ProjectGenerator(done, *projects...)
	buckets := listBuckets(done, pc)

	json.NewEncoder(w).Encode(buckets)
}
