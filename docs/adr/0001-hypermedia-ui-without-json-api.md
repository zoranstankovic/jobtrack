# Hypermedia UI with HTMX, without a JSON or GraphQL API

The UI is rendered on the server: Echo handlers return whole HTML pages or fragments (`templ` components), and HTMX swaps them into the page. There is no JSON or GraphQL API. HTMX expects HTML from the server. With a JSON/GraphQL layer, we would need client-side JavaScript to turn JSON into HTML, which is exactly the work HTMX removes, and the project would have two contracts to maintain instead of one.

## Consequences

If a separate client is needed (a mobile app, scripts), a JSON API is added next to the existing HTML handlers, reusing the same business layer. It does not replace them.
