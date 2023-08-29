package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M14 19a6 6 0 1 0 0 12a6 6 0 0 0 0-12zm0 10a4 4 0 1 1 0-8a4 4 0 0 1 0 8z"/><path fill="#55ACEE" d="M1.783 14.023v.02C.782 14.263 0 15.939 0 18s.782 3.737 1.783 3.956v.021l28.701 7.972V6.064L1.783 14.023z"/><ellipse cx="31" cy="18" fill="#269" rx="5" ry="12"/>`),
		g.Group(children),
	)
}