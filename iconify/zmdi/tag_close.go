package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M469 0q18 0 30.5 12.5T512 43v298q0 18-12.5 30.5T469 384H149q-21 0-34-19L0 192L115 19q13-19 34-19h320zm-64 269l-76-77l76-77l-30-30l-76 77l-77-77l-30 30l77 77l-77 77l30 30l77-77l76 77z"/>`),
		g.Group(children),
	)
}