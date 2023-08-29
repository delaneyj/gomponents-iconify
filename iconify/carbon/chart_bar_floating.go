package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarFloating(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 24H14v-8h14zm-12-2h10v-4H16zm10-10H8V4h18zm-16-2h14V6H10z"/><path fill="currentColor" d="M30 30H4a2.002 2.002 0 0 1-2-2V2h2v26h26Z"/>`),
		g.Group(children),
	)
}