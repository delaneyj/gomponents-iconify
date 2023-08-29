package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="6.5" cy="6.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="m14 7l-5.223 8.487L7 13l-5 7h20z"/>`),
		g.Group(children),
	)
}