package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelRoomTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 18.5V4m0 14.5s0 3-1.5 3m1.5-3h21m0 0s0 3 1.5 3m-1.5-3V13A3.5 3.5 0 0 0 19 9.5h-8.5v4m-9 2h21m-14-4.1S7.5 13 6.25 13a1.75 1.75 0 0 1-1.75-1.75c0-.966.784-1.746 1.75-1.746C7.5 9.504 8.5 11.1 8.5 11.1v.3Z"/>`),
		g.Group(children),
	)
}