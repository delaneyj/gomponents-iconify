package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuralPost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M25 12a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-8Zm9 3h-2v-2h-2v2h-2v2h2v2h2v-2h2v-2Z" clip-rule="evenodd"/><path d="M5 12h18v6H8v19h21V25h8v12h3V18h-1v-6h4v2h-1v23h1a1 1 0 1 1 0 2H5a1 1 0 1 1 0-2h1V14H5v-2Z"/><path d="M27 34H10v-2h1v-7h15v7h1v2Z"/></g>`),
		g.Group(children),
	)
}