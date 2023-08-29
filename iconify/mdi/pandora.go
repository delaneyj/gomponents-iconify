package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pandora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20a1 1 0 0 1-1 1H4V3h9.71a6.75 6.75 0 0 1 6.75 6.75c0 3.75-3.02 6.75-6.75 6.75H10V20Z"/>`),
		g.Group(children),
	)
}