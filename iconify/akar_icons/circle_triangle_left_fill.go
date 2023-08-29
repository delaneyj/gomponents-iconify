package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTriangleLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1Zm2 7.8a.809.809 0 0 0-.396-.705a.71.71 0 0 0-.77.04l-4.5 3.2A.815.815 0 0 0 8 12c0 .268.125.517.334.666l4.5 3.2a.71.71 0 0 0 .77.04A.809.809 0 0 0 14 15.2V8.8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}