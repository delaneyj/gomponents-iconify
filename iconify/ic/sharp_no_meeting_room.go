package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpNoMeetingRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 6h3v7.88l2 2V4h-5V3H6.12L14 10.88zm-2 5.71V13h-2v-2h1.29L2.41 2.13L1 3.54l4 4V19H3v2h11v-4.46L20.46 23l1.41-1.41z"/>`),
		g.Group(children),
	)
}