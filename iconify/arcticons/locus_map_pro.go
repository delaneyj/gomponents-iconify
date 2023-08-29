package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocusMapPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24h7.899v1.297a4 4 0 1 0 3 0V24H24m0 18.5v-7.899h1.297a4 4 0 1 0 0-3H24V24m18.5 0h-7.899v-1.297a4 4 0 1 0-3 0V24H24m0-18.5v7.899h-1.297a4 4 0 1 0 0 3H24V24"/>`),
		g.Group(children),
	)
}