package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAustria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EEE" d="M0 13h36v10H0z"/><path fill="#ED2939" d="M32 5H4a4 4 0 0 0-4 4v4h36V9a4 4 0 0 0-4-4zM4 31h28a4 4 0 0 0 4-4v-4H0v4a4 4 0 0 0 4 4z"/>`),
		g.Group(children),
	)
}