package Response

type DeploymentVoRes struct {
	Name      string
	Namespace string
}

func (d *DeploymentVoRes) GetName() string {
	return d.Name
}

func (d *DeploymentVoRes) SetName(name string) {
	d.Name = name
}
