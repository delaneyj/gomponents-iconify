package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16v4.46L13.537 12L20 17.54V22H4v-4.46L10.463 12L4 6.46V2Zm12.297 5L18 5.54V4H6v1.54L7.703 7h8.594ZM12 13.317L6 18.46V20h1l5-3l5 3h1v-1.54l-6-5.143Z"/>`),
		g.Group(children),
	)
}