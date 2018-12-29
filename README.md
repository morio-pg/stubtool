stubtool
===
Stubを作れるツール

## 環境

- Heroku
- Firebase
- Go 1.11.4
    - [gin](https://github.com/gin-gonic/gin)
    - [Wire](https://github.com/google/wire)
- Vue.js

### 本番設定

```bash
$ heroku config:set GIN_MODE="release"
$ heroku config:set FIREBASE_KEYFILE_JSON="$(< ./serviceAccountKey.json)"
```
