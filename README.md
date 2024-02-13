# caddy-pkl-adapter

Caddy pkl config adapter

## Install

Install with [xcaddy](https://github.com/caddyserver/xcaddy).

```
xcaddy build \
    --with github.com/armadi1809/caddy-pkl-adapter
```

## Usage

Specify with the `--adapter` flag for `caddy run`.

```
caddy run --config /path/to/pkl/config.pkl --adapter pkl
```
