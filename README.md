## 本番環境へリクエストする
https://yuchami-tinder-app.fly.dev/ に対してデプロイしています。<br>
こちらにリクエストを送ってください。<br>

**例：remindItemListを作成する (No.2)**<br>
`POST https://yuchami-tinder-app.fly.dev/manager/remindItemLists`
```request body
{
    "name": "test",
    "status": "ok",
    "is_delete": false,
    "remind_items":
    [
        {
            "order": 1,
            "url": "http://test.com",
            "status": "item1",
            "is_delete": false
        },
        {
            "order": 2,
            "url": "http://test.com",
            "status": "item2",
            "is_delete": false
        }
    ]
}
```
IDや~Atは自動生成のため、フロント側では考えなくて大丈夫です。

## ローカルでテストする
本レポジトリをcloneしたのち、下記で実行できます。<br>
1. app/配下に.envファイルを作成し、環境変数 `DATABASE_URL` を設定する
2. rootディレクトリに移動し、`docker-compose up --build` を実行する
3. `go run yuchami-tinder-app` を実行する
4. localhostに対し、リクエストを送る<br>
例：`http://localhost:8080/manager/remindItemLists`

なお、dbの接続情報等は `compose.yaml` に記載しています

## API詳細
[チームFのmiro](https://miro.com/app/board/uXjVKWxJRZI=/) を参照ください