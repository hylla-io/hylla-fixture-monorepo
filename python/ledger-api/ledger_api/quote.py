from ledger_core.rates import apply_rate


def quote(cents: int) -> str:
    """Quotes an amount with 20% added."""
    return f"quote: {apply_rate(cents, 20)}"
