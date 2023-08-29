package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2ZM19.72 18.15v10m3.56 0h5m-5-10h5m-5 5h3.26m-3.26-5v10"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.85 28.15v-10h3.27a3.36 3.36 0 0 1 0 6.72h-3.27m3.27 0l3.28 3.28m-28.87-10h6.62m-3.31 10v-10"/>`),
		g.Group(children),
	)
}