package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BenchBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5c-.55 0-1 .45-1 1v4c0 .55.45 1 1 1h1v2H1v2h2v4h2v-4h14v4h2v-4h2v-2h-4v-2h1c.55 0 1-.45 1-1V6c0-.55-.45-1-1-1H4m13 6v2H7v-2h10Z"/>`),
		g.Group(children),
	)
}