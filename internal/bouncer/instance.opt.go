package bouncer

import (
	"log"
	"os/user"
	"sync"

	"github.com/google/uuid"
)

var getDefaultOpt = sync.OnceValue[instanceOptions](func() instanceOptions {
	return instanceOptions{
		name:   getDefaultCurrentUsername(),
		logger: log.Default(),
	}
})

type instanceOptions struct {
	name   string
	logger *log.Logger
}

// NewInstanceOption TODO: docs
type NewInstanceOption interface {
	apply(*instanceOptions)
}

type funcDialOption struct {
	f func(*instanceOptions)
}

func (f *funcDialOption) apply(io *instanceOptions) {
	f.f(io)
}

// WithName configure instance name
func WithName(name string) NewInstanceOption {
	return newFuncInstanceOption(func(io *instanceOptions) {
		io.name = name
	})
}

// WithLogger configure instance logger
func WithLogger(l *log.Logger) NewInstanceOption {
	return newFuncInstanceOption(func(io *instanceOptions) {
		io.logger = l
	})
}

func newFuncInstanceOption(f func(*instanceOptions)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}

func getDefaultCurrentUsername() string {
	user, err := user.Current()
	if err != nil {
		log.Printf("cannot get current os user: %v", err)
		return uuid.NewString()
	}

	return user.Username
}
