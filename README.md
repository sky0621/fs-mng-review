# fs-mng-review

## 機能
### コンテンツ審査
- プレオーダー
- オーダー

### 審査機関(facility)の管理
- ラインナップ提供 [gRPC]

### レコード(議事録)管理
- オーダー毎の全レコード提供 [gRPC]

## 技術
### golang
```bash
$ go version
go version go1.13.9 linux/amd64
```

### Auth0

### GCP
#### Cloud Pub/Sub

### frontend
#### yarn
```
$ yarn -v
1.22.4
```

#### create
```
$ yarn create nuxt-app nuxt-app

create-nuxt-app v2.15.0
✨  Generating Nuxt.js project in nuxt-app
? Project name nuxt-app
? Project description My splendid Nuxt.js project
? Author name sky0621
? Choose programming language TypeScript
? Choose the package manager Yarn
? Choose UI framework Vuetify.js
? Choose custom server framework None (Recommended)
? Choose the runtime for TypeScript @nuxt/typescript-runtime
? Choose Nuxt.js modules Axios, DotEnv
? Choose linting tools ESLint, Prettier
? Choose test framework Jest
? Choose rendering mode Single Page App
? Choose development tools (Press <space> to select, <a> to toggle all, <i> to invert selection)

```