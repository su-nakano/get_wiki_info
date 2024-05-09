# get_wiki_info

- set API_KEY at .env

```
API_KEY="YOUR_KEY"
```

- run main.go

```
$ go run main.go
```

and response would be as follows.
```
$ go run .
Response: [{"id":3505287,"projectId":197969,"name":"MTG/Scrum/Sprint Plan - Review/2024/Sprint-078","content":"# OS_SYSTEM-2674 【仕様確定】リニューアルPh3：印鑑申し込みページ\r\n目標：\r\n- 商品詳細のコンポーネント実装が完了する\r\n- カート機能のフローの実装に入る\r\n　- （都度決済のコンポーネントと重複している部分もある）\r\n\r\n# OS_SYSTEM-2301 【仕様未定】会議室予約システム (見積：55Points - 消化済：61Points)\r\n目標：\r\n- 「リリース後でも対応可」の事業部から依頼を対応する\r\n\r\n# OS_SYSTEM-2517 【仕様確定】都度転送機能の追加（見積：21points - 消化済：5point）\r\n目標：\r\n（60％）\r\n- マイペー ジ側の実装を検証２面にマージする\r\n  - 決済画面（UI）→完了\r\n  - Thanks画面 →完了\r\n- BO側に着手する\r\n  - 発送仮登録（追跡番号追加、レタパ指定）\r\n  - 発送仮登録審査\r\n\r\n\r\n# OS_SYSTEM-1903 【仕様確定】不在票/本人限定郵便等の送付時に料金徴求（見積：13points- 消化済：10points）\r\n目標：\r\n- 検証2面で動作確認を完了する\r\n  - 通知文言のFix次第\r\n\r\n# OS_SYSTEM-2642 【仕様 未定】顧客対応用チャットボットの開発（消化済：１points）\r\n目標：\r\n- ECRにイメージをpushして、形態素解析が行えることを確認する\r\n\r\n# OS_SYSTEM-3188 【仕様未定】未収金回収後の自動メールに分岐を持たせたい\r\n目標：\r\n- 検証2面で動作確認を完了する\r\n  - 通知文言のFix次第\r\n\r\n# OS_SYSTEM-3232 【改善】リグレッションテスト用の資料を作成する\r\n目標：\r\n- 利用申込のテストケース作成の進捗 20%\r\n　-　個人申込（転送なし、オプションなし、クーポンなし）のテストケースを作成して、宮田さんに共有する\r\n- 郵便物登録のテストケース作成の進捗 10％\r\n　- 郵便物登録〜発送仮登録までのテストケース（普通郵便）を宮田さんに共有する\r\n\r\n# OS_SYSTEM-3111 【仕様確定】写真でお知らせオプションの追加申し込み通知（事業部向け）\r\n目標：\r\n- リリースできる（会議室予約をリリースした後にリリースする）\r\n\r\n# OS_SYSTEM-3237 [API 改善] 共通で使っている郵便物一覧取得のAPI(getPostItemList)を機能ごとに分割する\r\n目標：\r\n- リリースできる（会議室予約をリリースした後にリリースする）\r\n\r\n# OS_SYSTEM-3259  本番リリースのCI/CDを用意する\r\n目標：\r\n- バックオフィスのcode pipeline構築\r\n\r\n# OS_SYSTEM-2673 【仕様確定】リニューアルPh3：マネーフォワード申し込みページ (消化済：2Points)\r\n目標：\r\n- 実装に着手する\r\n  - 案内ページの実装が完了する\r\n\r\n# OS_SYSTEM-3282 【月次バッチの請求関連機能】BOからCSVをImportした時、決済日時が登録されない\r\n目標：\r\n- リリースできる\r\n  (１、２月分の請求データを対象にしてlambdaを実施しても結果を確かめることはできそう)\r\n\r\n↓（プランニング以降に追加した項目）↓\r\n# OS_SYSTEM-3290 【改善】BOから審査NGにする操作をしたとき、仮売上がキャンセルされない場合がある\r\n# OS_SYSTEM-3294 [DB] 会議室予約の割引期間を伸ばす\r\n\r\n## 運用\r\n\r\n","tags":[],"attachments":[],"sharedFiles":[],"stars":[],"createdUser":
Extracted Strings:
 OS_SYSTEM-2674 【仕様確定】リニューアルPh3：印鑑申し込みページ
 OS_SYSTEM-2301 【仕様未定】会議室予約システム (見積：55Points - 消化済：61Points)
 OS_SYSTEM-2517 【仕様確定】都度転送機能の追加（見積：21points - 消化済：5point）
 OS_SYSTEM-1903 【仕様確定】不在票/本人限定郵便等の送付時に料金徴求（見積：13points- 消化済：10points）
 OS_SYSTEM-2642 【仕様 未定】顧客対応用チャットボットの開発（消化済：１points）
 OS_SYSTEM-3188 【仕様未定】未収金回収後の自動メールに分岐を持たせたい
 OS_SYSTEM-3232 【改善】リグレッションテスト用の資料を作成する
 OS_SYSTEM-3111 【仕様確定】写真でお知らせオプションの追加申し込み通知（事業部向け）
 OS_SYSTEM-3237 [API 改善] 共通で使っている郵便物一覧取得のAPI(getPostItemList)を機能ごとに分割する
 OS_SYSTEM-3259  本番リリースのCI/CDを用意する
 OS_SYSTEM-2673 【仕様確定】リニューアルPh3：マネーフォワード申し込みページ (消化済：2Points)
 OS_SYSTEM-3282 【月次バッチの請求関連機能】BOからCSVをImportした時、決済日時が登録されない
 OS_SYSTEM-3290 【改善】BOから審査NGにする操作をしたとき、仮売上がキャンセルされない場合がある
 OS_SYSTEM-3294 [DB] 会議室予約の割引期間を伸ばす
# 運用
```
