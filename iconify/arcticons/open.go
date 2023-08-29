package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Open(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36 41.84A33.32 33.32 0 0 1 13 5.55"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 23.65a33.21 33.21 0 0 1 43 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.1 5.79a34.7 34.7 0 0 1 .27 3.84A33.21 33.21 0 0 1 12 41.82"/>`),
		g.Group(children),
	)
}