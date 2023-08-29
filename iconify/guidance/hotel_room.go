package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 18.5V14l.676-.184a37.335 37.335 0 0 1 19.648 0L22.5 14v4.5m-21 0s0 3-1.5 3m1.5-3h21m0 0s0 3 1.5 3M3.5 11c0-1.989-.297-3.966-.882-5.867L2.5 4.75V4.5h19v.25l-.118.383A19.95 19.95 0 0 0 20.5 11M12 7.5H6.5V11M12 7.5V11m0-3.5h5.5V11"/>`),
		g.Group(children),
	)
}