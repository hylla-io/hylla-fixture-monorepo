//! Rounds amounts in cents.

/// Rounds an amount to whole cents, half up.
pub fn round_cents(x: f64) -> i64 {
    (x + 0.5).floor() as i64
}
