package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"

	"simpleApp/mylib"
)

type Page struct {
	Title         string
	Body          []byte
	SearchResults map[string]int
	Query         string
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var templates = template.Must(template.ParseFiles("view.html", "edit.html", "search.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /view/test
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /edit/test
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func searchPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /search/test
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	query := r.URL.Query().Get("query")
	if query != "" {
		// 解析するテキスト
		text := string(p.Body)

		// テキストを形態素解析してトークンに分割
		words := mylib.TokenizeText(text)

		// 特定の単語の出現回数をカウント
		targetWord := query

		// 方法1: 形態素列としてフレーズを検索
		count1 := mylib.CountPhraseFrequency(words, targetWord)
		fmt.Printf("\n方法1: 形態素列として「%s」の出現回数: %d回\n", targetWord, count1)

		// 方法2: 元のテキストで直接フレーズを検索
		count2 := mylib.CountPhraseInOriginalText(text, targetWord)
		fmt.Printf("方法2: 元テキストで「%s」の出現回数: %d回\n", targetWord, count2)

		p.Query = query
		// 検索結果を追加
		p.SearchResults = map[string]int{
			fmt.Sprintf("「%s」の形態素解析による出現回数", query): count1,
			fmt.Sprintf("「%s」の直接検索による出現回数", query):  count2,
		}
	}
	renderTemplate(w, "search", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /save/test
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view|search)/([a-zA=Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/search/", makeHandler(searchPageHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
