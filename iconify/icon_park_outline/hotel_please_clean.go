package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelPleaseClean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M13 27s20 0 20-12v29H13V27Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 15v-1c0-5.523-4.477-10-10-10S13 8.477 13 14m20 1c0 12-20 12-20 12v17h20V15Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m19 34l3 3l5-5"/></g>`),
		g.Group(children),
	)
}