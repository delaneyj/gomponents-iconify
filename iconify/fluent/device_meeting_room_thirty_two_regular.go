package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceMeetingRoomThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.47 4a4.25 4.25 0 0 0-4.096 3.113L2.154 18.7c-.738 2.665 1.265 5.3 4.03 5.3h19.634c2.765 0 4.768-2.635 4.028-5.299l-3.22-11.588A4.25 4.25 0 0 0 22.53 4H9.47ZM7.302 7.648A2.25 2.25 0 0 1 9.47 6h13.06a2.25 2.25 0 0 1 2.168 1.648l3.22 11.588a2.18 2.18 0 0 1-2.1 2.764H6.182a2.18 2.18 0 0 1-2.1-2.764L7.301 7.649ZM7.999 26a1 1 0 1 0 0 2h16.002a1 1 0 1 0 0-2H7.999Z"/>`),
		g.Group(children),
	)
}