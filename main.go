package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const daily_report = "/home/hiroki_togashi/obsidian/daily_report/"

func main() {
	now := time.Now()
	today := getFilePath(now)
	yesterday := getFilePath(now.AddDate(0, 0, -1))

	// コピー元のファイルを開く
	source, err := os.Open(yesterday)
	if err != nil {
		fmt.Println("コピー元のファイルを開く際にエラーが発生しました:", err)
		return
	}
	defer source.Close()

	// コピー先のファイルを作成する
	destination, err := os.Create(today)
	if err != nil {
		fmt.Println("コピー先のファイルを作成する際にエラーが発生しました:", err)
		return
	}
	defer destination.Close()

	// ファイルの内容を1行ずつ読み込み、[x]で始まる行を排除
	scanner := bufio.NewScanner(source)
	writer := bufio.NewWriter(destination)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") || strings.Contains(line, "- [ ]") {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				fmt.Println("ファイルへの書き込み中にエラーが発生しました:", err)
				return
			}
		}
	}

	// エラーチェック
	if err := scanner.Err(); err != nil {
		fmt.Println("ファイルを読み込む際にエラーが発生しました:", err)
		return
	}

	// バッファをフラッシュしてファイルに書き込み
	writer.Flush()

	fmt.Println("ファイル名: " + today)
}

func getFilePath(date time.Time) string {
	return daily_report + date.Format("2006-01-02") + ".md"
}
