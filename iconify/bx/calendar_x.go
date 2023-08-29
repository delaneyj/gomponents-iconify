package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.293 16.293l1.414 1.414L12 15.414l2.293 2.293l1.414-1.414L13.414 14l2.293-2.293l-1.414-1.414L12 12.586l-2.293-2.293l-1.414 1.414L10.586 14z"/><path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2zm.002 16H5V8h14l.002 12z"/>`),
		g.Group(children),
	)
}