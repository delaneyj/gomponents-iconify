package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 6.6h29.09v22.09a7 7 0 0 1-7 7H12.5a7 7 0 0 1-7-7V6.6h0Zm29.09 0h5.91a2 2 0 0 1 2 2v5.7a2 2 0 0 1-2 2h-5.91h0V6.6h0ZM5.5 41.4h29.09"/>`),
		g.Group(children),
	)
}