package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowArcRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 88v64a8 8 0 0 1-8 8h-64a8 8 0 0 1 0-16h44.6l-22.36-22.21A88 88 0 0 0 40 184a8 8 0 0 1-16 0a104 104 0 0 1 177.54-73.54L224 132.77V88a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}