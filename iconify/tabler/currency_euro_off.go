package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEuroOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.2 7c-1.977-2.26-4.954-2.602-7.234-1.04M8.053 8.039c-1.604 2.72-1.374 6.469.69 8.894c2.292 2.691 6 2.758 8.356.18M10 10H5m0 4h8M3 3l18 18"/>`),
		g.Group(children),
	)
}