package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplayThirtyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12c0-5.523-4.477-10-10-10a9.985 9.985 0 0 0-8 4V3.5H2v6h4.75v.5h2.625a.625.625 0 1 1 0 1.25H7.5v1.5h1.875a.625.625 0 1 1 0 1.25H6.75v1.5h2.625a2.125 2.125 0 0 0 1.62-3.5a2.125 2.125 0 0 0-1.62-3.5H8v-1H5.385A8 8 0 1 1 4 12H2c0 5.523 4.477 10 10 10s10-4.477 10-10Zm-9.5-1.25a2.5 2.5 0 0 1 5 0v2.5a2.5 2.5 0 0 1-5 0v-2.5Zm2.5-1a1 1 0 0 0-1 1v2.5a1 1 0 1 0 2 0v-2.5a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}