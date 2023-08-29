package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriorityHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 19h8v-2h-8v2m0-5.5h8v-2h-8v2M14 8h8V6h-8v2M2 12.5C2 8.92 4.92 6 8.5 6H9V4l3 3l-3 3V8h-.5C6 8 4 10 4 12.5S6 17 8.5 17H12v2H8.5C4.92 19 2 16.08 2 12.5Z"/>`),
		g.Group(children),
	)
}