package Response

type DeploymentVoRes struct {
	Name []string
}

func (d *DeploymentVoRes) GetName() []string {
	return d.Name
}

func (d *DeploymentVoRes) SetName(name []string) {
	d.Name = name
}

func (d *DeploymentVoRes) AddName(name string) {
	d.Name = append(d.Name, name)
}
