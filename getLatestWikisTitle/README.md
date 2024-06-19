# get_wiki_info

- set API_KEY at .env

```
API_KEY="YOUR_KEY"
```

- run main.go

```
$ go run main.go
```

then, you can see output files in the directory with `Sprint-{sprint number}.txt`
it contains notification text and wiki url as follows.
```
@os_jigyoubu @gr-os-developer @ @

sp-082は12ポイントを消化しました。
sp-83(06/11(木) - 06/10(水)は次を対応します:
 https://gmo-office.backlog.com/alias/wiki/3584688
 リリース予定： OS_SYSTEM-3336 [API] 決済リンクの有効期限を伸ばす方法があるか調査する
 要件定義：OS_SYSTEM-3232 【改善】リグレッションテスト用の資料を作成する
 実装：OS_SYSTEM-2674 【仕様確定】リニューアルPh3：印鑑申し込みページ (消化済：2Points)
 実装：OS_SYSTEM-2517 【仕様確定】都度転送機能の追加（見積：21points - 消化済：8point）（６/７に判断予定）
 実装：OS_SYSTEM-1903 【仕様確定】不在票/本人限定郵便等の送付時に料金徴求（見積：13points- 消化済：10points）
 実装：OS_SYSTEM-2642 【仕様 未定】顧客対応用チャットボットの開発（消化済：2points）
 実装：OS_SYSTEM-3259  本番リリースのCI/CDを用意する（消化済：1points）
 実装：OS_SYSTEM-2673 【仕様確定】リニューアルPh3：マネーフォワード申し込みページ (消化済：4Points)
 実装：OS_SYSTEM-3330 【仕様未定】発送登録省力化のためのBO改修
 実装：OS_SYSTEM-3378 【改善】storage_id_viewの取得方法を変える

```
