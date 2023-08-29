package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinistryOfHealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23.375 5.22a1 1 0 0 1 1.25 0L39.35 17H42v6H6v-6h2.65L23.374 5.22ZM25 10v2h2v2h-2v2h-2v-2h-2v-2h2v-2h2Zm-1 15a4 4 0 0 0-4 4v8h-6v-4h-2v-8H8v8H6v4H4v6h40v-6h-2v-4h-2v-8h-4v8h-2v4h-6v-8a4 4 0 0 0-4-4Zm0 2a2 2 0 0 1 2 2v8h-4v-8a2 2 0 0 1 2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}