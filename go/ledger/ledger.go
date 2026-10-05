// Package ledger posts amounts in cents.
package ledger

import "fmt"

// Post prints a posted amount.
func Post(cents int64) string {
	return fmt.Sprintf("posted %d", cents)
}
