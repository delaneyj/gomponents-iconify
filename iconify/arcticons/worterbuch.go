package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Worterbuch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.066 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.297 5.5v16.328l5.321-2.661l5.321 2.661V5.5M11.075 22.56c0-5.291 11.75-5.69 11.75.799v11.963a2.287 2.287 0 0 0 2.287 2.287h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.825 26.953c-23.963-1.198-9.816 19.567 0 6.09"/>`),
		g.Group(children),
	)
}