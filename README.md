# meshi API

MESHIBUGYO のサーバーサイドを管理するリポジトリ

# requirements

- Go 1.16.5
- Docker 20.10.7
- docker-compose 1.29.2
- direnv 2.28.0

# 開発環境の立ち上げ

```
direnv allow
docker-compose up
```

# 開発に必要なツールのインストール

```
make install-devdeps
```

# アーキテクチャ

- GraphQL & DDD & 半 Clean Architecture  
  [Creating GraphQL-Apps using Go](https://medium.com/@alexander.zimmermann96/creating-graphql-apps-using-go-7c9a12ad6b42)  
  この人と思想一緒や
- GraphQL と DDD の対応は  
  [gqlgen のフィールド解決方法の使い分け tips](https://qiita.com/ryota-yamamoto/items/3f15f476f17db047ef5d)

  > そもそも gqlgen の良さは GralphQL のスキーマを先に決めて、それに答えるバックエンドを作る「スキーマファースト開発」ができる点だと考えています。スキーマからリゾルバーの作成を gqlgen に任せ、リゾルバーがアプリケーションのユースケースを呼び出すことで、アプリケーション内部に GraphQL という概念が入り込まないようにしています。

  この考え方が理解できるため、自動バインドで解決している
