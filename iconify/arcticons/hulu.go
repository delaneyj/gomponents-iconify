package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hulu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.06 22v5a3 3 0 0 0 3 3h0a3 3 0 0 0 3-3v-5m7.37 0v5a3 3 0 0 0 3 3h0a3 3 0 0 0 3-3v-5M9.5 17.93v12.14m0-5.01a3 3 0 0 1 3-3h0a3 3 0 0 1 3 3v5m13.23-12.13v12.14"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2H7.48Z"/>`),
		g.Group(children),
	)
}