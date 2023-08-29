package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowArcRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 88v64a6 6 0 0 1-6 6h-64a6 6 0 0 1 0-12h49.45l-25.8-25.63A90 90 0 0 0 38 184a6 6 0 0 1-12 0a102 102 0 0 1 174.12-72.12l25.88 25.7V88a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}