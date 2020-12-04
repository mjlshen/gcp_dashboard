package gcloud

type Bucket struct {
	Id                string `json:"id"`
	ProjectId         string `json:"projectId"`
	Name              string `json:"name"`
	Location          string `json:"location"`
	StorageClass      string `json:"storageClass"`
	VersioningEnabled bool   `json:"versioningEnabled"`
}

type GKECluster struct {
	Id                   string `json:"id"`
	ProjectId            string `json:"projectId"`
	Name                 string `json:"name"`
	Location             string `json:"location"`
	CurrentMasterVersion string `json:"currentMasterVersion"`
	CurrentNodeVersion   string `json:"currentNodeVersion"`
}

type Project struct {
	ProjectId string `json:"projectId"`
	Number    int64  `json:"id"`
}
