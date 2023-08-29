package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 12h6m-6 6H8A6 6 0 0 1 8 6h1m6 0h1a6 6 0 0 1 0 12h-1"/>`),
		g.Group(children),
	)
}