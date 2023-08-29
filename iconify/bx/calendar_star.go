package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.981 14.811l-.467 2.726l2.449-1.287l2.449 1.287l-.468-2.726l1.982-1.932l-2.738-.398L11.963 10l-1.225 2.481L8 12.879z"/><path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2zm.002 16H5V8h14l.002 12z"/>`),
		g.Group(children),
	)
}