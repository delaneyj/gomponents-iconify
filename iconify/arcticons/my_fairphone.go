package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyFairphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.355 9.5H33.03M16.75 24h11.425M18.53 9.5l-3.56 29m25.53-33h-33a2.25 2.25 0 0 0-2 2v33a2.25 2.25 0 0 0 2 2h33a2.25 2.25 0 0 0 2-2v-33a2.25 2.25 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}