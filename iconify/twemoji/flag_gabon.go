package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagGabon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#009E60" d="M32 5H4a4 4 0 0 0-4 4v5h36V9a4 4 0 0 0-4-4z"/><path fill="#3A75C4" d="M0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-5H0v5z"/><path fill="#FCD116" d="M0 14h36v8H0z"/>`),
		g.Group(children),
	)
}