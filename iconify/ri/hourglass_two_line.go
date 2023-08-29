package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16v4.46L13.537 12L20 17.54V22H4v-4.46L10.463 12L4 6.46V2Zm8 8.683l6-5.143V4H6v1.54l6 5.143Zm0 2.634L6 18.46V20h12v-1.54l-6-5.143Z"/>`),
		g.Group(children),
	)
}