package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Define a struct to parse the JSON data
type Item struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"projectId"`
	Name      string `json:"name"`
	Content   string `json:"content"`
}

// Function to extract strings between # and \r
// it basically extracts the issue numbers from the content
func extractStrings(content string) []string {
	// Compile a regular expression to find patterns
	re := regexp.MustCompile(`#(.+?)\r`)
	// Find all matches
	matches := re.FindAllStringSubmatch(content, -1)

	var results []string
	for _, match := range matches {
		if len(match) > 1 { // match[0] is the full match, match[1] is the first group
			results = append(results, match[1])
		}
	}
	return results
}

// it fetches wiki pages from the given project and keywords in gmoos
// Example JSON output
// jsonData := `[
//
//	{
//	    "id": 3505287,
//	    "projectId": 197969,
//	    "name": "MTG/Scrum/Sprint Plan - Review/2024/Sprint-078",
//	    "content":"# OS_SYSTEM-2674 【仕様確定】リニューアルPh3：印鑑申し込みページ\r\n目標：\r\n- 商品詳細のコンポーネント実装が完了する\r\n- カート機能のフローの実装に入る\r\n　- （都度決済のコンポーネントと重複している部分もある）\r\n\r\n# OS_SYSTEM-2301 【仕様未定】会議室予約システム (見積：55Points - 消化済：61Points)\r\n目標：\r\n- 「リリース後でも対応可」の事業部から依頼を対応する\r\n\r\n# OS_SYSTEM-2517 【仕様確定】都度転送機能の追加（見積：21points - 消化済：5point）\r\n目標：\r\n（60％）\r\n- マイペー ジ側の実装を検証２面にマージする\r\n  - 決済画面（UI）→完了\r\n  - Thanks画面 →完了\r\n- BO側に着手する\r\n  - 発送仮登録（追跡番号追加、レタパ指定）\r\n  - 発送仮登録審査\r\n\r\n\r\n# OS_SYSTEM-1903 【仕様確定】不在票/本人限定郵便等の送付時に料金徴求（見積：13points- 消化済：10points）\r\n目標：\r\n- 検証2面で動作確認を完了する\r\n  - 通知文言のFix次第\r\n\r\n# OS_SYSTEM-2642 【仕様 未定】顧客対応用チャットボットの開発（消化済：１points）\r\n目標：\r\n- ECRにイメージをpushして、形態素解析が行えることを確認する\r\n\r\n# OS_SYSTEM-3188 【仕様未定】未収金回収後の自動メールに分岐を持たせたい\r\n目標：\r\n- 検証2面で動作確認を完了する\r\n  - 通知文言のFix次第\r\n\r\n# OS_SYSTEM-3232 【改善】リグレッションテスト用の資料を作成する\r\n目標：\r\n- 利用申込のテストケース作成の進捗 20%\r\n　-　個人申込（転送なし、オプションなし、クーポンなし）のテストケースを作成して、宮田さんに共有する\r\n- 郵便物登録のテストケース作成の進捗 10％\r\n　- 郵便物登録〜発送仮登録までのテストケース（普通郵便）を宮田さんに共有する\r\n\r\n# OS_SYSTEM-3111 【仕様確定】写真でお知らせオプションの追加申し込み通知（事業部向け）\r\n目標：\r\n- リリースできる（会議室予約をリリースした後にリリースする）\r\n\r\n# OS_SYSTEM-3237 [API 改善] 共通で使っている郵便物一覧取得のAPI(getPostItemList)を機能ごとに分割する\r\n目標：\r\n- リリースできる（会議室予約をリリースした後にリリースする）\r\n\r\n# OS_SYSTEM-3259  本番リリースのCI/CD��用意する\r\n目標：\r\n- バックオフィスのcode pipeline構築\r\n\r\n# OS_SYSTEM-2673 【仕様確定】リニューアルPh3：マネーフォワード申し込みページ (消化済：2Points)\r\n目標：\r\n- 実装に着手する\r\n  - 案内ページの実装が完了する\r\n\r\n# OS_SYSTEM-3282 【月次バッチの請求関連機能】BOからCSVをImportした時、決済日時が登録されない\r\n目標：\r\n- リリースできる\r\n  (１、２月分の請求データを対象にしてlambdaを実施しても結果を確かめることはできそう)\r\n\r\n↓（プランニング以降に追加した項目）↓\r\n# OS_SYSTEM-3290 【改善】BOから審査NGにする操作をしたとき、仮売上がキャンセルされない場合がある\r\n# OS_SYSTEM-3294 [DB] 会議室予約の割引期間を伸ばす\r\n\r\n## 運用\r\n\r\n"
//	        }
//
// ]j
func fetchWikiPages(apiKey, projectIdOrKey, keyword string) []byte {
	// Base URL of the API
	baseURL := "https://gmo-office.backlog.com/api/v2/wikis"

	// Prepare URL with query parameters
	params := url.Values{}
	params.Add("apiKey", apiKey)
	params.Add("projectIdOrKey", projectIdOrKey)
	params.Add("keyword", keyword)
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the HTTP GET request
	response, err := http.Get(fullURL)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	defer response.Body.Close() // Ensure we close the response body

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	// Print the response body to the console
	// fmt.Printf("Response: %s\n", body)
	return body
}

func topUpSlackInformation(consumedPoints string, prevSprint string, nextSprint string) string {
	now := time.Now()

	// Calculate how many days until this week's Thursday
	daysUntilThursday := (4 - int(now.Weekday()) + 7) % 7
	thisThursday := now.AddDate(0, 0, daysUntilThursday).Format("01/02")

	// Calculate the date for the next week's Wednesday
	nextWednesday := now.AddDate(0, 0, daysUntilThursday+6).Format("01/02")
	message := fmt.Sprintf("%sは%sポイントを消化しました。\n%s(%s(木) - %s(水))は次を対応します:", prevSprint, consumedPoints, nextSprint, thisThursday, nextWednesday)
	return message
}

// writeToFile writes the given strings to a text file
func writeToFile(fileName string, lines []string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	currentSprintNum := os.Getenv("CURRENT_SPRINT_NUM")
	prevSprintPoint := os.Getenv("PREV_SPRINT_POINT")
	projectIdOrKey := "197969"

	prevSprint := fmt.Sprintf("sp-%s", currentSprintNum)

	nextSprintNum, nextSprintNumErr := strconv.Atoi(currentSprintNum) // expected 082 not 82

	if nextSprintNumErr != nil {
		log.Fatalf("Error converting sprint number: %v", nextSprintNumErr)
	}
	nextSprint := fmt.Sprintf("sp-%d", nextSprintNum+1)
	notifyMember := "@os_jigyoubu @gr-os-developer @Yuta Horie/GMO-FH @Tsuyoshi Chiba/FH"
	keyword := fmt.Sprintf("Review/2024/Sprint-%s", currentSprintNum)
	wikiPages := fetchWikiPages(apiKey, projectIdOrKey, keyword)

	// Unmarshal the JSON data into an array of Items
	var items []Item
	if err := json.Unmarshal([]byte(wikiPages), &items); err != nil {
		panic(err) // Handle error appropriately in production code
	}

	// Assuming there is only one item in the JSON for this example
	if len(items) > 0 {
		content := items[0].Content
		extractedStrings := extractStrings(content)

		// Separate "リリース予定" lines and other lines
		var releaseLines, wikiUrl, planningLines, otherLines []string
		for _, line := range extractedStrings {
			if strings.Contains(line, "リリース予定") {
				fmt.Printf("リリース予定： %s\n", line)
				releaseLines = append(releaseLines, line)
			} else if strings.Contains(line, "https://") {
				wikiUrl = append(wikiUrl, line)
			} else if strings.Contains(line, "要件定義") {
				planningLines = append(planningLines, line)
			} else if strings.Contains(line, "# 運用") {

			} else {
				fmt.Printf("その他: %s\n", line)
				otherLines = append(otherLines, line)
			}
		}

		sortedStrings := append(wikiUrl, "  ")
		sortedStrings = append(sortedStrings, "```")
		sortedStrings = append(sortedStrings, releaseLines...)
		sortedStrings = append(sortedStrings, planningLines...)
		sortedStrings = append(sortedStrings, otherLines...)
		sortedStrings = append(sortedStrings, "```")
		// Print extracted strings
		staticText := []string{
			notifyMember,
			"",
			topUpSlackInformation(prevSprintPoint, prevSprint, nextSprint),
		}

		// 固定文言とソート文言をくっつける
		outputLines := append(staticText, sortedStrings...)
		fileName := fmt.Sprintf("Sprint-%s.txt", currentSprintNum)
		err := writeToFile(fileName, outputLines)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

		fmt.Println("File written successfully.")
	} else {
		fmt.Println("No items found in JSON")
	}
}
