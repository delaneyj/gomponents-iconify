package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2H7.48Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.32 12.12a4.47 4.47 0 1 1-4.47 4.47a4.47 4.47 0 0 1 4.47-4.47ZM20 21l7.22 7.22a.7.7 0 0 0 1 0l1.38-1.38a.71.71 0 0 1 1 0l7.81 7.81a.7.7 0 0 1-.5 1.2H10.08a.7.7 0 0 1-.58-1.1L19 21.13a.7.7 0 0 1 1-.13Z"/>`),
		g.Group(children),
	)
}