# email-search-engine

A full-text search UI for an indexed email corpus, backed by [ZincSearch](https://github.com/zincsearch/zincsearch). Companion to [indexer](https://github.com/rdaniel1105/indexer), which populates the ZincSearch index from a maildir-style tree.

Originally built around the [Enron email corpus](https://www.cs.cmu.edu/~enron/) (~500k messages) — once `indexer` has ingested the dataset, this app lets you query it with paginated results and per-email detail views.

## Architecture

```
indexer  ─────►  ZincSearch  ◄─────  back-end  ◄─────  front-end
(Go CLI)         (search DB)         (Go API)          (Vue 3 / TS / Tailwind)
```

- **`back-end/`** — A small Go HTTP service (chi router) that proxies search queries to ZincSearch's `/api/<index>/_search` endpoint, normalizes the response, and returns hits + pagination metadata to the UI.
- **`front-end/`** — A Vue 3 + TypeScript + Tailwind SPA with a search box, paginated results table, and an email detail panel. Built with Vue CLI Service.

The split keeps ZincSearch credentials out of the browser — the front-end never talks to the search DB directly.

## Requirements

- Go 1.26+
- Node.js 18+
- A running [ZincSearch](https://github.com/zincsearch/zincsearch) instance with an index populated by [indexer](https://github.com/rdaniel1105/indexer)

## Running locally

**1. Back-end**

```bash
cd back-end
cp .env.example .env   # fill in ZincSearch credentials
go run .
```

The server listens on `PORT` (default `3000`).

| Variable | Default | Purpose |
|---|---|---|
| `PORT` | `3000` | HTTP port for the API |
| `ZINCSEARCH_URL` | `http://localhost:4080` | Base URL of the ZincSearch instance |
| `ZINCSEARCH_INDEX` | `emails` | Target index name (must match the indexer's target) |
| `ZINCSEARCH_USERNAME` | `""` | ZincSearch admin user |
| `ZINCSEARCH_PASSWORD` | `""` | ZincSearch admin password |

**2. Front-end**

```bash
cd front-end
cp .env.example .env   # set VUE_APP_SERVER_URL to the back-end's URL
npm install
npm run serve
```

| Variable | Purpose |
|---|---|
| `VUE_APP_SERVER_URL` | Base URL of the back-end, e.g. `http://localhost:3000` |

Then open http://localhost:8080.

## Tests

```bash
# Back-end
cd back-end && go test ./...

# Front-end
cd front-end && npm test
```

## License

MIT — see [LICENSE](./LICENSE).
