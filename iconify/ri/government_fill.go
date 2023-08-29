package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GovernmentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19V8H1V6h3V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2h3v2h-1v11h1v2H1v-2h1Zm11 0v-7h-2v7h2Zm-5 0v-7H6v7h2Zm10 0v-7h-2v7h2ZM6 5v1h12V5H6Z"/>`),
		g.Group(children),
	)
}