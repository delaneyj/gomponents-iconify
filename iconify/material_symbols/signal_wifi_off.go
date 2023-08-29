package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.925 15.075L7.45 4.625q1.125-.3 2.25-.463T12 4q3.4 0 6.513 1.287T24 9l-6.075 6.075Zm2.55 8.225L15.1 17.9L12 21L0 9q.8-.8 1.663-1.475T3.474 6.3l-2.8-2.8L2.1 2.075l19.8 19.8l-1.425 1.425Z"/>`),
		g.Group(children),
	)
}