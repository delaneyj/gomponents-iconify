package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficConeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 210h-19.73L151.54 35.4a14 14 0 0 0-13.23-9.4h-20.62a14 14 0 0 0-13.23 9.4L43.73 210H24a6 6 0 0 0 0 12h208a6 6 0 0 0 0-12ZM94 102h68l18.08 52H75.92Zm21.8-62.66a2 2 0 0 1 1.89-1.34h20.62a2 2 0 0 1 1.89 1.34L157.82 90H98.18ZM71.74 166h112.52l15.3 44H56.44Z"/>`),
		g.Group(children),
	)
}