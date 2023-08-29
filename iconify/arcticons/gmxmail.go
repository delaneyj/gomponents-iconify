package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gmxmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.89 22.47a3.11 3.11 0 0 0-3.28-3.11a3.23 3.23 0 0 0-2.94 3.28v2.88a3.11 3.11 0 0 0 3.11 3.11h0a3.11 3.11 0 0 0 3.11-3.11h-3.11m6.5 3.11v-9.27l4.63 9.28l4.63-9.27v9.27m3.39-9.28L35.04 24l-3.11 4.64m6.22-9.28L35.04 24l3.11 4.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}