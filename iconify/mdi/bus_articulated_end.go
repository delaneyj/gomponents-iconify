package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusArticulatedEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 6L20 7.5L21.5 9L20 10.5l1.5 1.5l-1.5 1.5l1.5 1.5h-9a3 3 0 0 1-3 3a3 3 0 0 1-3-3h-4V8c0-1.11.89-2 2-2h17m-3 1.5H15V10h3.5V7.5m-5 0h-4V10h4V7.5M8 7.5H4V10h4V7.5m1.5 6A1.5 1.5 0 0 0 8 15a1.5 1.5 0 0 0 1.5 1.5A1.5 1.5 0 0 0 11 15a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}