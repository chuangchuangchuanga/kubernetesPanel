package Request

type DeploynetListVoReq struct {
	NamespaceName string `json:"namespaceName" binding:"required,min=1,max=50"`
}

func (d *DeploynetListVoReq) GetNamespace() string {
	return d.NamespaceName
}
