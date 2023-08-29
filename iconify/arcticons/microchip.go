package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microchip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.91 17.14h12.18a.76.76 0 0 1 .77.77v12.18a.76.76 0 0 1-.77.77H17.91a.76.76 0 0 1-.77-.77V17.91a.76.76 0 0 1 .77-.77Z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.97 13.26V2.59m-3.94 10.67V2.59m-3.94 10.67V3.32m11.82 9.94V3.32m-7.88 31.51V45.5m3.94-10.67V45.5m3.94-10.67v9.94m-11.82-9.94v9.94m16.69-18.75h10.67m-10.67-3.94h10.67m-10.67-3.94h9.94m-9.94 11.81h9.94m-31.5-7.87H2.55m10.67 3.94H2.55m10.67 3.93H3.28m9.94-11.81H3.28"/>`),
		g.Group(children),
	)
}