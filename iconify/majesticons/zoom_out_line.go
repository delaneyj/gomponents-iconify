package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 11a7 7 0 1 1 14 0a7 7 0 0 1-14 0zm7-9a9 9 0 1 0 5.618 16.032l3.675 3.675a1 1 0 0 0 1.414-1.414l-3.675-3.675A9 9 0 0 0 11 2zm-3 8a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H8z"/></g>`),
		g.Group(children),
	)
}