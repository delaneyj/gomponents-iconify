package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSubLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.5 12.5l-5 5m0 0l-5-5m5 5V9.7c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874c.428-.218.988-.218 2.108-.218h9.8"/>`),
		g.Group(children),
	)
}