package templates

type Template interface {
	Render() error
	Save() error
}
