package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTriangleUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1ZM8.8 14a.809.809 0 0 1-.705-.396a.71.71 0 0 1 .04-.77l3.2-4.5A.815.815 0 0 1 12 8c.268 0 .517.125.666.334l3.2 4.5a.71.71 0 0 1 .04.77a.809.809 0 0 1-.706.396H8.8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}