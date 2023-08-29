package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegisteredFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm.5 5H8v10h2v-3h2.217l2.18 3h2.472l-2.55-3.51a3.5 3.5 0 0 0-1.627-6.486L12.5 7Zm0 2a1.5 1.5 0 0 1 1.493 1.355L14 10.5l-.007.145a1.5 1.5 0 0 1-1.348 1.348L12.5 12H10V9h2.5Z"/>`),
		g.Group(children),
	)
}