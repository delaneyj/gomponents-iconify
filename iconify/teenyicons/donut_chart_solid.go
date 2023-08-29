package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DonutChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 7.5A7.5 7.5 0 0 1 7 .016v4.02a3.5 3.5 0 1 0 2.596 6.267l2.842 2.842A7.5 7.5 0 0 1 0 7.5Z"/><path fill="currentColor" d="M13.145 12.438A7.471 7.471 0 0 0 15 7.5c0-1.034-.21-2.018-.587-2.914L10.755 6.21a3.498 3.498 0 0 1-.452 3.385l2.842 2.842ZM8 4.035V.016a7.499 7.499 0 0 1 5.963 3.676L10.254 5.34A3.497 3.497 0 0 0 8 4.035Z"/>`),
		g.Group(children),
	)
}