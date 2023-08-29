package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m234.83 98.83l-96 96a4 4 0 0 1-5.66 0L28 89.66V152a4 4 0 0 1-8 0V80a4 4 0 0 1 4-4h72a4 4 0 0 1 0 8H33.66L136 186.34l93.17-93.17a4 4 0 1 1 5.66 5.66Z"/>`),
		g.Group(children),
	)
}