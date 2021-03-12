# ciao

Very-simple & idiomatic HTTP redirect server.

## How to use it?

Create a `ciao.json` config file where you want, with your redirections:

```json
{
  "redirects": {
    "example.com": {
      "location": "https://example.org",
      "code": 307
    },
    "www.example.com": {
      "location": "https://example.org",
      "code": 308
    }
  }
}
```

and then execute ciao: `./ciao --config <path-to-config>`