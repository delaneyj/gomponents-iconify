package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V3h10v4h10v8h-2V9h-8v2h2v2h-2v2h2v2h-2v2h4v2H2Zm2-2h2v-2H4v2Zm0-4h2v-2H4v2Zm0-4h2V9H4v2Zm0-4h2V5H4v2Zm4 12h2v-2H8v2Zm0-4h2v-2H8v2Zm0-4h2V9H8v2Zm0-4h2V5H8v2Zm12 16v-2h-2v-2h2v-2h2v2h2v2h-2v2h-2Zm-4-10v-2h2v2h-2Zm0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}