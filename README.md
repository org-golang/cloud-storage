# CloudStorage

## Application configure
```yaml
providers:
  alibaba:
    appid: alibaba-app-id
    secret: alibaba-app-secret
    callback: https://alibaba.com
    base-url: https://alibaba.com
    scopes:
      - user:base
      - file:all:read
  baidu:
    appid: baidu-app-id
    app-key: baidu-app-key
    secret-key: baidu-app-secret
    sign-key: baidu-sing-key
    display: popup
    callback: https://baidu.com
    base-url: https://openapi.baidu.com
    scopes:
      - basic
      - netdisk
```
