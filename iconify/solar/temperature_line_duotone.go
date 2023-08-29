package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 22a5.5 5.5 0 0 0 3.439-9.793a1.114 1.114 0 0 1-.439-.86V5a3 3 0 1 0-6 0v6.348c0 .338-.175.648-.439.86A5.5 5.5 0 0 0 12 22Z" opacity=".5"/><path stroke-linecap="round" d="M12 14a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Zm0 0V5"/></g>`),
		g.Group(children),
	)
}