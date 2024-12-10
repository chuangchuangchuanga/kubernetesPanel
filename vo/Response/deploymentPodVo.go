package Response

type DeploymentPodVoRes struct {
	Name []string
}

func (d *DeploymentPodVoRes) GetName() []string {
	return d.Name
}

func (d *DeploymentPodVoRes) SetName(name []string) {
	d.Name = name
}

func (d *DeploymentPodVoRes) AddName(name string) {
	d.Name = append(d.Name, name)
}
