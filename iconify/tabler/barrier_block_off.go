package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarrierBlockOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 7h8a1 1 0 0 1 1 1v7c0 .27-.107.516-.282.696M16 16H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h2m0 9v4m.5-4l4.244-4.244m2.001-2.001L16.5 7m-3 9l1.249-1.249m1.992-1.992L20 9.5m-16 4l4.752-4.752M17 17v3M5 20h4m6 0h4M17 7V5M3 3l18 18"/>`),
		g.Group(children),
	)
}