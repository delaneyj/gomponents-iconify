package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesTurtleneck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 23v14m-26 0v7h26v-7m-26 0H4V23c0-3 2-6.5 5-9s9-4 9-4h12s6 1.5 9 4s5 6 5 9v14h-7m-26 0V23m19-13H18V4h12v6Z"/>`),
		g.Group(children),
	)
}