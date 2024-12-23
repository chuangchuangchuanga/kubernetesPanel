package Response

type DeploymentPodVoRes struct {
	Name       string
	Status     string
	CreateTime string
}

func (d *DeploymentPodVoRes) GetName() string {
	return d.Name
}

func (d *DeploymentPodVoRes) SetName(name string) {
	d.Name = name
}
