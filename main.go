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
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #f9f9f9;
            padding: 50px;
            text-align: center; /* タイトルを中央揃え */
        }
        h2 {
            color: #4a90e2;
        }
        ul {
            list-style-type: none;
            padding: 0;
        }
        li {
            margin: 10px 0;
            background-color: #e6e6e6;
            border-radius: 5px;
            padding: 5px 15px;
            transition: background-color 0.3s;
            cursor: pointer; /* カーソルをポインタに */
        }
        li:hover {
            background-color: #d4d4d4;
        }
        a {
            display: block;  /* リストの全体がクリック可能に */
            height: 100%;
            width: 100%;
            text-decoration: none;
            color: #333;
            transition: color 0.3s;
        }
        a:hover {
            color: #4a90e2;
        }
    </style>
</head>
<body>

<h1>PDF・HTMLスライドビューア</h1>

<h2>スライドを選択してください</h2>

<ul>
    {{ range . }}
        <li><a href="/files/{{ . }}">{{ . }}</a></li>
    {{ end }}
</ul>

</body>
<footer>
	<p><small>作成: gomadoufu</small></p>
</footer>
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
