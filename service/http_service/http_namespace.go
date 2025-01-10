package http_service

// Namespace 是名字空间，可以理解为一个服务组
type Namespace struct {
	path        string
	description string
	schemas     map[string]*Schema[any, any]
}

// NewNamespace 创建一个新的 Namespace 实例
func NewNamespace(path string, description string) *Namespace {
	return &Namespace{
		path:        path,
		description: description,
		schemas:     make(map[string]*Schema[any, any]),
	}
}

func (n *Namespace) AppendSchema(name string, schema *Schema[any, any]) {
	n.schemas[name] = schema
}

func (n *Namespace) GetPath() string {
	return n.path
}

func (n *Namespace) GetDescription() string {
	return n.description
}

func (n *Namespace) GetSchemasLen() int {
	return len(n.schemas)
}

func (n *Namespace) GetSchema(name string) *Schema[any, any] {
	return n.schemas[name]
}

func (n *Namespace) AddSchema(name string, schema *Schema[any, any]) {
	n.schemas[name] = schema
}
