package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixlr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 1 2.5 24A21.51 21.51 0 0 1 24 2.5Zm0 7.09A14.41 14.41 0 0 0 9.59 24h0A14.41 14.41 0 0 0 24 38.41h0A14.41 14.41 0 0 0 38.41 24h0A14.41 14.41 0 0 0 24 9.59Zm0 7.59l5.91 3.41v6.82L24 30.82l-5.91-3.41v-6.82Z"/>`),
		g.Group(children),
	)
}