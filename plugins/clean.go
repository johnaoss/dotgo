package plugins

// The name of the directive that Clean handles
const _clean = "clean"

// TODO: Make clean behave properly.
type clean struct{}

func newClean() Plugin {
	return new(clean)
}

func (c *clean) CanHandle(directive string) bool {
	return directive == _clean
}

// todo
func (c *clean) Handle(directive string, data map[string]interface{}) error {
	return nil
}
