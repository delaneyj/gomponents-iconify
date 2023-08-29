package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 20h40M6 6v8m0 11.996v12.003M20.4 6v8m0 12v16M34.8 6v8M42 6v8m-7.2 12v8M13.2 6v8m0 12v8M27.6 6v8m0 12v8M42 26v12"/>`),
		g.Group(children),
	)
}