package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 56a6 6 0 0 1-6 6h-35.68l-44 132H144a6 6 0 0 1 0 12H64a6 6 0 0 1 0-12h35.68l44-132H112a6 6 0 0 1 0-12h80a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}