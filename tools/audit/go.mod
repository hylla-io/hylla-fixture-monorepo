module github.com/hylla-io/hylla-fixture-monorepo/tools/audit

go 1.24.0

require (
	github.com/google/go-cmp v0.7.0
	github.com/hylla-io/hylla-fixture-monorepo/go/ledger v0.0.0
)

replace github.com/hylla-io/hylla-fixture-monorepo/go/ledger => ../../go/ledger
