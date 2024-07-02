// Package janitor TODO: docs
package janitor

import "log"

// CheckClose TODO: docs
func CheckClose(fn func() error, name string) {
	log.Printf("closing %s", name)
	if err := fn(); err != nil {
		log.Fatalf("failed to close %s: %v", name, err)
	}
}
