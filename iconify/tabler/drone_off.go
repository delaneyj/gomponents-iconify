package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 14h-4v-4m0 0L6.5 6.5m3.457-.55A3.503 3.503 0 0 0 7.04 3.04m-3.02.989A3.5 3.5 0 0 0 6 9.965M14 10l3.5-3.5m.5 3.465A3.5 3.5 0 1 0 14.034 6M14 14l3.5 3.5m-3.465.5a3.5 3.5 0 0 0 5.936 1.98m.987-3.026a3.503 3.503 0 0 0-2.918-2.913M10 14l-3.5 3.5M6 14.035A3.5 3.5 0 1 0 9.966 18M3 3l18 18"/>`),
		g.Group(children),
	)
}