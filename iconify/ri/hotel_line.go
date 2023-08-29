package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 21H2v-2h1V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v5h2v10h1v2Zm-5-2h2v-8h-6v8h2v-6h2v6Zm0-10V5H5v14h6V9h6ZM7 11h2v2H7v-2Zm0 4h2v2H7v-2Zm0-8h2v2H7V7Z"/>`),
		g.Group(children),
	)
}