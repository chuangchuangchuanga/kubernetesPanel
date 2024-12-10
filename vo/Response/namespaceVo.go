package Response

type NamespaceVo struct {
	Name []string
}

func (n *NamespaceVo) GetName() []string {
	return n.Name
}

func (n *NamespaceVo) SetName(name []string) {
	n.Name = name
}

func (n *NamespaceVo) AddName(name string) {
	n.Name = append(n.Name, name)
}
