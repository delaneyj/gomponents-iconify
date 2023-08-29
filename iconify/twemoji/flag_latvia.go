package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagLatvia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#9E3039" d="M32 5H4a4 4 0 0 0-4 4v6h36V9a4 4 0 0 0-4-4zm0 26H4a4 4 0 0 1-4-4v-6h36v6a4 4 0 0 1-4 4z"/><path fill="#EEE" d="M0 15h36v6H0z"/>`),
		g.Group(children),
	)
}