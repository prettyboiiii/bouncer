package bouncer

type instanceOptions struct {
	name string
}

// NewInstanceOption TODO: docs
type NewInstanceOption interface {
	apply(*instanceOptions)
}

type funcDialOption struct {
	f func(*instanceOptions)
}

// apply implements NewInstanceOption.
func (f *funcDialOption) apply(io *instanceOptions) {
	f.f(io)
}

// WithName configure instance name
func WithName(name string) NewInstanceOption {
	return newFuncInstanceOption(func(io *instanceOptions) {
		io.name = name
	})
}

func newFuncInstanceOption(f func(*instanceOptions)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}
