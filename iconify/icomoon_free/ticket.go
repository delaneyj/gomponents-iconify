package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m9 5l2 2l-4 4l-2-2zm6.649-.351L14.5 3.5L14 4a1.414 1.414 0 1 1-1.999-2l.5-.5L11.352.351a1.208 1.208 0 0 0-1.703 0L.352 9.648a1.208 1.208 0 0 0 0 1.703L1.501 12.5L2 12.001A1.414 1.414 0 1 1 4 14l-.5.5l1.149 1.149a1.208 1.208 0 0 0 1.703 0l9.297-9.297a1.208 1.208 0 0 0 0-1.703zM7 13L3 9l6-6l4 4l-6 6z"/>`),
		g.Group(children),
	)
}