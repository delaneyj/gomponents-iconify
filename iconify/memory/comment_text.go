package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1V2m1 2v12h5v3h1v-1h1v-1h1v-1h8V4H3m2 3h12v2H5V7m0 4h10v2H5v-2Z"/>`),
		g.Group(children),
	)
}