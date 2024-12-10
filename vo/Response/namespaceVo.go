package Response

type NamespaceVoRes struct {
	Name []string
}

func (n *NamespaceVoRes) GetName() []string {
	return n.Name
}

func (n *NamespaceVoRes) SetName(name []string) {
	n.Name = name
}

func (n *NamespaceVoRes) AddName(name string) {
	n.Name = append(n.Name, name)
}
