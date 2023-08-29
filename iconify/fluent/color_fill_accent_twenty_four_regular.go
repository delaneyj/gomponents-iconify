package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorFillAccentTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.052 15c.273 1.61 1.58 3 3.448 3c1.842 0 3.14-1.354 3.436-2.936A1.5 1.5 0 0 1 21 16.5v3a1.5 1.5 0 0 1-1.5 1.5h-15A1.5 1.5 0 0 1 3 19.5v-3A1.5 1.5 0 0 1 4.5 15h3.09a3 3 0 0 0 4.048 0h1.414Z"/>`),
		g.Group(children),
	)
}