package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m237.66 101.66l-96 96a8 8 0 0 1-11.32 0L32 99.31V152a8 8 0 0 1-16 0V80a8 8 0 0 1 8-8h72a8 8 0 0 1 0 16H43.31L136 180.69l90.34-90.35a8 8 0 0 1 11.32 11.32Z"/>`),
		g.Group(children),
	)
}