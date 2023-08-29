package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Formulia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.102 35.915v-23.83m-11.188 23.83v-23.83m0 23.83H6.674a2.17 2.17 0 0 1-1.64-3.592L12.245 24l-7.21-8.323a2.17 2.17 0 0 1 1.64-3.592H43.5"/>`),
		g.Group(children),
	)
}