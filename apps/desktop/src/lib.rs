//! The desktop app's labels.

/// Labels a total for the window title.
pub fn total_label(x: f64) -> String {
    format!("total: {}", mono_ledger::round_cents(x))
}
