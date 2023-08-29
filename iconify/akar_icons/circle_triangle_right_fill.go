package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleTriangleRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1Zm-2 7.8c0-.295.152-.566.396-.705a.71.71 0 0 1 .77.04l4.5 3.2A.815.815 0 0 1 16 12a.815.815 0 0 1-.334.666l-4.5 3.2a.71.71 0 0 1-.77.04A.809.809 0 0 1 10 15.2V8.8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}