package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commercial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 6H1a11.431 11.431 0 0 1 1-4h11a11.429 11.429 0 0 1 1 4zM3 7h9v6h-1V8H8v5H3zm1 3h3V8H4z"/>`),
		g.Group(children),
	)
}