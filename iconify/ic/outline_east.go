package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 5l-1.41 1.41L18.17 11H2v2h16.17l-4.59 4.59L15 19l7-7l-7-7z"/>`),
		g.Group(children),
	)
}