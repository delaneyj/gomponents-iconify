package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookClockTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.75 11.375a.625.625 0 0 0 .688.622c.02.002.04.003.062.003h1.25a.625.625 0 1 0 0-1.25H13V9.125a.625.625 0 1 0-1.25 0v2.25ZM4 4.5A2.5 2.5 0 0 1 6.5 2H18a2.5 2.5 0 0 1 2.5 2.5v14.25a.75.75 0 0 1-.75.75H5.5a1 1 0 0 0 1 1h13.25a.75.75 0 0 1 0 1.5H6.5A2.5 2.5 0 0 1 4 19.5v-15Zm8.375 1.505a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5Z"/>`),
		g.Group(children),
	)
}