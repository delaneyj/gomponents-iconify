package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 120a44 44 0 1 1 44-44a44.05 44.05 0 0 1-44 44Zm60 8a44 44 0 1 0 44 44a44.05 44.05 0 0 0-44-44Zm-120 0a44 44 0 1 0 44 44a44.05 44.05 0 0 0-44-44Z"/>`),
		g.Group(children),
	)
}