package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 18v2H2v-2h20ZM2 3.5l8 5l-8 5v-10ZM22 11v2H12v-2h10ZM4 7.109v2.783L6.226 8.5L4 7.109ZM22 4v2H12V4h10Z"/>`),
		g.Group(children),
	)
}