package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareShadowTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2a2 2 0 0 0 2 2h3.5A2.5 2.5 0 0 0 11 8.5V5a2 2 0 0 0-2-2a2 2 0 0 0-2-2H3Zm-.5 2a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3Z"/>`),
		g.Group(children),
	)
}