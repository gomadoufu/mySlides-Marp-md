package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ファイル選択</title>
</head>
<body>

<h2>ファイルを選択してください。</h2>

<ul>
    {{ range . }}
        <li><a href="/files/{{ . }}">{{ . }}</a></li>
    {{ end }}
</ul>

</body>
</html>
`

func listFiles(w http.ResponseWriter, r *http.Request) {
	// *.pdf と *.html のファイルを取得
	pdfFiles, _ := filepath.Glob("*.pdf")
	htmlFiles, _ := filepath.Glob("*.html")

	// 2つのスライスを結合
	files := append(pdfFiles, htmlFiles...)

	// テンプレートを解析
	tmpl, err := template.New("fileList").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	// テンプレートを実行してHTTPレスポンスに書き込む
	if err := tmpl.Execute(w, files); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", listFiles)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %s\n", err)
		os.Exit(1)
	}
}
