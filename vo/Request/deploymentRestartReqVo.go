package Request

type DeploymentRestartVoReq struct {
	NamespaceName  string `json:"namespaceName" binding:"required,min=1,max=50"`
	DeploymentName string `json:"deploymentName" binding:"required,min=1,max=50"`
}

func (d *DeploymentRestartVoReq) GetNamespace() string {
	return d.NamespaceName
}

func (d *DeploymentRestartVoReq) GetDeploymentName() string {
	return d.DeploymentName
}
