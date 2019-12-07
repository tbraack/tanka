package manifest

// An Identifier can be used to uniquely identify a Manifest
type Identifier struct {
	Kind      string
	Name      string
	Namespace string
}

func (m Manifest) Identifier() Identifier {
	return Identifier{
		Kind:      m.Kind(),
		Name:      m.Metadata().Name(),
		Namespace: m.Metadata().Namespace(),
	}
}
