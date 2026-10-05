import { formatCents } from "@mono/money";

/** Shows a price in dollars. */
export function Price({ cents }: { cents: number }) {
  return <span className="price">{formatCents(cents)}</span>;
}
