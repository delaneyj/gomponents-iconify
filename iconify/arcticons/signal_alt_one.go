package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.933 2.527C12.075 2.55 2.479 12.18 2.5 24.038a21.471 21.471 0 0 0 3.33 11.448l-1.65 7.107a1 1 0 0 0 1.197 1.201l7.18-1.644c10.025 6.334 23.287 3.343 29.621-6.682c6.335-10.025 3.343-23.287-6.682-29.62a21.471 21.471 0 0 0-11.562-3.32Z"/>`),
		g.Group(children),
	)
}