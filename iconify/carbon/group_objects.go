package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupObjects(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 13V9H9v10h4v4h10V13Zm-8-2h6v6h-6Zm10 10h-6v-2h4v-4h2Z"/><path fill="currentColor" d="M30 8V2h-6v2H8V2H2v6h2v16H2v6h6v-2h16v2h6v-6h-2V8Zm-4-4h2v2h-2ZM4 4h2v2H4Zm2 24H4v-2h2Zm22 0h-2v-2h2Zm-2-4h-2v2H8v-2H6V8h2V6h16v2h2Z"/>`),
		g.Group(children),
	)
}