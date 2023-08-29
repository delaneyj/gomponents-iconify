package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceMeetingRoomTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.096 3a2.5 2.5 0 0 0-2.438 1.946l-1.59 7A2.5 2.5 0 0 0 4.504 15h10.989a2.5 2.5 0 0 0 2.438-3.054l-1.591-7A2.5 2.5 0 0 0 13.903 3H6.096ZM4.633 5.168A1.5 1.5 0 0 1 6.096 4h7.807a1.5 1.5 0 0 1 1.463 1.168l1.59 7A1.5 1.5 0 0 1 15.495 14H4.505a1.5 1.5 0 0 1-1.463-1.832l1.591-7ZM5.5 16a.5.5 0 0 0 0 1h9a.5.5 0 1 0 0-1h-9Z"/>`),
		g.Group(children),
	)
}