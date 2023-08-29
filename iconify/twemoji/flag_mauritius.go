package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMauritius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EA2839" d="M32 5H4a4 4 0 0 0-4 4v2.5h36V9a4 4 0 0 0-4-4z"/><path fill="#1A206D" d="M0 11.5h36V18H0z"/><path fill="#FFD500" d="M0 18h36v6.5H0z"/><path fill="#00A551" d="M0 24.5V27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-2.5H0z"/>`),
		g.Group(children),
	)
}