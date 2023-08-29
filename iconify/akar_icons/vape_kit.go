package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VapeKit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10a1 1 0 0 1 1-1h6v11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V10Zm10 5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v6h-8v-6Zm2-6h4v3h-4zM6 4h4v5H6zm2 0V2m9 7V6m0 11v-5"/>`),
		g.Group(children),
	)
}