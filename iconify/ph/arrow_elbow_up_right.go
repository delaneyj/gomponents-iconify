package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m213.66 85.66l-48 48a8 8 0 0 1-11.32-11.32L188.69 88H72v136a8 8 0 0 1-16 0V80a8 8 0 0 1 8-8h124.69l-34.35-34.34a8 8 0 0 1 11.32-11.32l48 48a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}