package mylib

import (
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// TokenizeText テキストを形態素解析してトークンに分割する
func TokenizeText(text string) []string {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}

	tokens := t.Tokenize(text)
	var words []string

	for _, token := range tokens {
		// 表層形（実際の単語）を取得
		word := token.Surface
		// 品詞情報を取得
		features := token.Features()

		// 空白や記号は除外（オプション）
		if len(features) > 0 && features[0] != "記号" && word != "" && word != " " {
			words = append(words, word)
		}
	}

	return words
}

// CountWordFrequency 指定した単語の出現回数をカウントする
func CountWordFrequency(words []string, targetWord string) int {
	count := 0
	for _, word := range words {
		if word == targetWord {
			count++
		}
	}
	return count
}

// CountPhraseFrequency 指定したフレーズ（複数の形態素で構成される語句）の出現回数をカウントする
func CountPhraseFrequency(words []string, targetPhrase string) int {
	// フレーズを形態素解析して単語列に分割
	phraseWords := TokenizeText(targetPhrase)

	// フレーズが形態素解析されなかった場合（1単語の場合など）
	if len(phraseWords) == 0 {
		return CountWordFrequency(words, targetPhrase)
	}

	count := 0
	// - 元のテキストから得られた単語リスト(`words`)の中で、フレーズの長さ分のウィンドウをスライドさせながら検索します
	// - `i`は検索開始位置を表し、最大でも「単語リストの長さ - フレーズの単語数」までしか進めません
	for i := 0; i <= len(words)-len(phraseWords); i++ {
		// 1. **連続する形態素の一致確認**
		matched := true
		// - 現在の位置`i`から始まる連続する形態素が、フレーズの形態素と完全に一致するか確認します
		// - 例えば「猫である」を検索する場合：
		//    - `words[i]` が「猫」と一致するか
		//    - `words[i+1]` が「である」と一致するか
		for j := 0; j < len(phraseWords); j++ {
			if words[i+j] != phraseWords[j] {
				// - 一つでも不一致があれば、`matched = false`として次の開始位置に移動します
				matched = false
				break
			}
		}
		if matched {
			// - すべての形態素が一致した場合（`matched`が`true`のまま）、カウンターをインクリメントします
			count++
		}
	}
	return count
}

// GetAllWordFrequencies すべての単語の出現回数を辞書形式で返す
func GetAllWordFrequencies(words []string) map[string]int {
	frequencies := make(map[string]int)
	for _, word := range words {
		frequencies[word]++
	}
	return frequencies
}

// CountPhraseInOriginalText 元のテキストで直接フレーズを検索
func CountPhraseInOriginalText(text, targetPhrase string) int {
	return strings.Count(text, targetPhrase)
}
