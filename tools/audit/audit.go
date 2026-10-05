// Package audit checks posted amounts against the ledger.
package audit

import (
	"github.com/google/go-cmp/cmp"
	"github.com/hylla-io/hylla-fixture-monorepo/go/ledger"
)

// Matches reports whether the ledger posts cents as want.
func Matches(cents int64, want string) bool {
	return cmp.Equal(ledger.Post(cents), want)
}
