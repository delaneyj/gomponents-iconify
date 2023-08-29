package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GovernmentLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6h3v2h-1v11h1v2H1v-2h1V8H1V6h3V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2Zm0 2H4v11h3v-7h2v7h2v-7h2v7h2v-7h2v7h3V8ZM6 5v1h12V5H6Z"/>`),
		g.Group(children),
	)
}