package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 19h2v-8h-6v8h2v-6h2v6ZM3 19V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v5h2v10h1v2H2v-2h1Zm4-8v2h2v-2H7Zm0 4v2h2v-2H7Zm0-8v2h2V7H7Z"/>`),
		g.Group(children),
	)
}