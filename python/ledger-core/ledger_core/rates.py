def apply_rate(cents: int, rate_percent: int) -> int:
    """Adds `rate_percent` to an amount in cents, rounding half up."""
    return (cents * (100 + rate_percent) + 50) // 100
