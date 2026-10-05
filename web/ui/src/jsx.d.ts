// The elements this package renders, so TSX type-checks without a UI library dependency.
declare namespace JSX {
  type Element = unknown;
  interface IntrinsicElements {
    span: { className?: string; children?: unknown };
  }
}
