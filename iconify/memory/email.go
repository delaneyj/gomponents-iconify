package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Email(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 5h1V4h18v1h1v13h-1v1H2v-1H1V5m2 12h16V9h-1v1h-2v1h-2v1h-2v1h-2v-1H8v-1H6v-1H4V9H3v8M19 6H3v1h2v1h2v1h2v1h4V9h2V8h2V7h2V6Z"/>`),
		g.Group(children),
	)
}