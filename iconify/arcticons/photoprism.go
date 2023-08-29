package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photoprism(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 36.059L24 5.912l20.5 30.147Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 5.912l6.612 3.617l12.682 32.56l-36.176-2.413L3.5 36.06m3.618 3.616L30.612 9.53m12.682 32.558l1.206-6.03"/>`),
		g.Group(children),
	)
}