package Request

type DeploynetPodListVoReq struct {
	NamespaceName  string `json:"namespaceName" binding:"required,min=1,max=50"`
	DeploymentName string `json:"deploymentName" binding:"required,min=1,max=50"`
}

func (d *DeploynetPodListVoReq) GetNamespace() string {
	return d.NamespaceName
}

func (d *DeploynetPodListVoReq) GetDeploymentName() string {
	return d.DeploymentName
}
