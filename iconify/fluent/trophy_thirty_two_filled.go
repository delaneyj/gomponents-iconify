package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 2a3 3 0 0 0-3 3H6.5A2.5 2.5 0 0 0 4 7.5v4a4.5 4.5 0 0 0 4.25 4.493A8.008 8.008 0 0 0 15 21.938V24h-3.5A3.5 3.5 0 0 0 8 27.5v1A1.5 1.5 0 0 0 9.5 30h13a1.5 1.5 0 0 0 1.5-1.5v-1a3.5 3.5 0 0 0-3.5-3.5H17v-2.062a8.008 8.008 0 0 0 6.75-5.945A4.5 4.5 0 0 0 28 11.5v-4A2.5 2.5 0 0 0 25.5 5H24a3 3 0 0 0-3-3H11Zm13 11.95V7h1.5a.5.5 0 0 1 .5.5v4a2.5 2.5 0 0 1-2 2.45ZM8 7v6.95a2.5 2.5 0 0 1-2-2.45v-4a.5.5 0 0 1 .5-.5H8Z"/>`),
		g.Group(children),
	)
}