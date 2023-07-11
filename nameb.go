// name パッケージは識別子名の操作を提供します.
package nameb

import (
	"strings"

	"github.com/17e10/go-rob"
)

// ctype は文字タイプを表します.
type ctype int

const (
	cDelim = ctype(iota) // 空白文字
	cUpper               // 大文字
	cLower               // 小文字
	cOther               // その他
)

// extractWords は s から単語を抽出します.
func extractWords(s string, fn func(word string)) {
	l := len(s)
	first := 0
	lastType := cDelim
	for i := 0; i < l; i++ {
		c := s[i]

		ctype := cOther
		switch {
		case c == ' ' || c == '_' || c == '-' || c == '.':
			ctype = cDelim
		case 'A' <= c && c <= 'Z':
			ctype = cUpper
		case 'a' <= c && c <= 'z':
			ctype = cLower
		}

		if lastType == cDelim {
			if ctype == cDelim {
				continue
			} else {
				first = i
			}
		} else if ctype == cDelim {
			// 単語の区切り
			fn(s[first:i])
		} else if lastType == cUpper && ctype == cLower {
			// 大文字 → 小文字 ... Abc の b を検出, first は A を指す
			if first < i-1 {
				fn(s[first : i-1])
			}
			first = i - 1
		}

		if ctype != cOther {
			lastType = ctype
		}
	}
	if first != l-1 && lastType != cDelim {
		fn(s[first:])
	}
}

// transform は名前変換します.
//
// wordfn は単語を変換する関数, sep はセパレータ文字を指定します.
// sep = 0 のとき セパレータを挟みません.
func transform(s string, sep byte, wordfn func(i int, s string) string) string {
	if s == "" {
		return ""
	}

	b := strings.Builder{}
	b.Grow(len(s) + 5)
	i := 0
	extractWords(s, func(word string) {
		// セパレータを追加する
		if sep != 0 && i > 0 {
			b.WriteByte(sep)
		}
		// word を変換して追加する
		b.WriteString(wordfn(i, word))
		i++
	})
	return b.String()
}

// Snake は snake_case へ名前変換します.
func Snake(s string) string {
	return transform(s, '_', func(i int, s string) string {
		return strings.ToLower(s)
	})
}

// Kebab は Kebab-case へ名前変換します.
func Kebab(s string) string {
	return transform(s, '-', func(i int, s string) string {
		return strings.ToLower(s)
	})
}

// Camel は camelCase へ名前変換します.
func Camel(s string) string {
	return transform(s, 0, func(i int, s string) string {
		b := []byte(strings.ToLower(s))
		if i > 0 && 'a' <= b[0] && b[0] <= 'z' {
			b[0] -= 'a' - 'A'
		}
		return rob.String(b)
	})
}

// Pascal は PascalCase へ名前変換します.
func Pascal(s string) string {
	return transform(s, 0, func(i int, s string) string {
		b := []byte(strings.ToLower(s))
		if 'a' <= b[0] && b[0] <= 'z' {
			b[0] -= 'a' - 'A'
		}
		return rob.String(b)
	})
}
