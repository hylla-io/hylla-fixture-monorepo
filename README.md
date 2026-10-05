# hylla-fixture-monorepo

A small repository with eight parts in four languages, used as a Hylla ingest fixture.

## Parts

- Rust: a root Cargo workspace with the ledger crate, and a desktop app crate that is its own
  Cargo workspace two folders down.
- Go: a ledger module listed in a Go workspace file, and an audit module with no workspace file.
- TypeScript: a pnpm workspace with a money package and a UI package written in TSX.
- Python: a core project and an API project that depends on it by path.

## How the parts connect

In each language, one part calls a function in another part:

- the desktop app labels a total with the ledger crate's rounding;
- the audit module checks what the ledger module posts;
- the UI package shows a price with the money package's formatting;
- the API project quotes an amount with the core project's rate.
