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

type ServiceAccount struct {
	Id        string `json:"uniqueId"`
	ProjectId string `json:"projectID"`
	Email     string `json:"email"`
	Disabled  bool   `json:"disabled"`
}

type IAMPolicyBinding struct {
	Role    string   `json:"role"`
	Members []string `json:"members"`
}

type IAMPolicy struct {
	ProjectId         string             `json:"projectID"`
	IAMPolicyBindings []IAMPolicyBinding `json:"iamPolicyBindings"`
}

type Project struct {
	ProjectId string `json:"projectId"`
	Number    int64  `json:"id"`
}
