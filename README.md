# ciao

Very-simple & idiomatic HTTP redirect server.

## How to use it?

Create a `ciao.json` config file where you want, with your redirects:

```json
{
  "redirects": {
    "example.com": "https://example.org",
    "www.example.com": "https://example.org"
  }
}
```

and then execute ciao: `./ciao --config <path-to-config>`