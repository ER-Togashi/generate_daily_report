# generate_daily_report

日報のフォーマットを自動生成するスクリプト

cronの設定をしておけば定期的に実行されて楽です

```bash
echo '0 6 * * 1-5 cd /path & go run main.go' | crontab -
```