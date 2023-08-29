package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blockbuster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.835 24h-5.504m0-11.5h4.83a3.833 3.833 0 0 1 3.833 3.833v3.834A3.833 3.833 0 0 1 24.161 24h.674a3.833 3.833 0 0 1 3.833 3.833v3.834a3.833 3.833 0 0 1-3.833 3.833h-5.504m0-23v23"/>`),
		g.Group(children),
	)
}