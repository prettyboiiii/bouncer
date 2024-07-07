// Package janitor TODO: docs
package janitor

import (
	"io"
	"log"
)

// CheckClose TODO: docs
func CheckClose(closer io.Closer, name string) {
	log.Printf("closing %s", name)
	if err := closer.Close(); err != nil {
		log.Fatalf("failed to close %s: %v", name, err)
	}
}
