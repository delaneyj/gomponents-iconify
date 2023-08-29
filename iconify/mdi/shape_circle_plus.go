package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeCirclePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19a6 6 0 0 0 6-6h2a8 8 0 0 1-8 8a8 8 0 0 1-8-8a8 8 0 0 1 8-8v2a6 6 0 0 0-6 6a6 6 0 0 0 6 6m8-14h3v2h-3v3h-2V7h-3V5h3V2h2v3Z"/>`),
		g.Group(children),
	)
}