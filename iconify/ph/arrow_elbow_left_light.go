package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m236.24 100.24l-96 96a6 6 0 0 1-8.48 0L30 94.48V152a6 6 0 0 1-12 0V80a6 6 0 0 1 6-6h72a6 6 0 0 1 0 12H38.48L136 183.51l91.76-91.75a6 6 0 0 1 8.48 8.48Z"/>`),
		g.Group(children),
	)
}