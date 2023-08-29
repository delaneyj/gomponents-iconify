package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceMeetingRoomSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.739 3.459A2.001 2.001 0 0 1 4.665 2h6.67c.896 0 1.683.596 1.926 1.459l1.66 5.898A2.08 2.08 0 0 1 12.917 12H3.083a2.08 2.08 0 0 1-2.004-2.643l1.66-5.898ZM4.497 13a.5.5 0 1 0 0 1h7.005a.5.5 0 1 0 0-1H4.497Z"/>`),
		g.Group(children),
	)
}