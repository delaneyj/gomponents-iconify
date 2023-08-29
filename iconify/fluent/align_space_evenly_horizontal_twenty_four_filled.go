package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignSpaceEvenlyHorizontalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 17.5A2.5 2.5 0 0 0 4.5 20h1A2.5 2.5 0 0 0 8 17.5v-11A2.5 2.5 0 0 0 5.5 4h-1A2.5 2.5 0 0 0 2 6.5v11Zm7 0a2.5 2.5 0 0 0 2.5 2.5h1a2.5 2.5 0 0 0 2.5-2.5v-11A2.5 2.5 0 0 0 12.5 4h-1A2.5 2.5 0 0 0 9 6.5v11Zm7 0a2.5 2.5 0 0 0 2.5 2.5h1a2.5 2.5 0 0 0 2.5-2.5v-11A2.5 2.5 0 0 0 19.5 4h-1A2.5 2.5 0 0 0 16 6.5v11Z"/>`),
		g.Group(children),
	)
}