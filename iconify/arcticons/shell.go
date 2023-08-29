package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 8.4a2 2 0 0 0-2 2v27.2a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2V10.4a2 2 0 0 0-2-2ZM9 23.16l7.72 6L9 35.13v-12Zm9.46 8.66h11v3.31h-11Z"/>`),
		g.Group(children),
	)
}