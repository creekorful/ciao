# direktion

Very-simple & idiomatic HTTP URL rewriter.

## How to use it?

Create a `direktion.json` config file where you want, with your redirections:

```json
{
  "redirects": {
    "example.com/blog": {
      "location": "https://blog.example.org/$request_uri",
      "code": 307
    },
    "example.com": {
      "location": "https://example.org",
      "code": 307
    },
    "www.example.com": {
      "location": "https://example.org",
      "code": 308
    }
  },
  "use_x_forwarded": false
}
```

and then execute direktion: `./direktion --config <path-to-config>`

nb: if `use_x_forwarded` is true, then direktion will use the `X-Forwarded*` headers to determinate remote IP address. (
should only enabled if behind trusted reverse proxy)