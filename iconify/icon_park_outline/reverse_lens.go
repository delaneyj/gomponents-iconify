package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReverseLens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M13 25V10h6l2-3h6l2 3h6v15H13Zm7 10l4 4l-4 4"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 38.168C39.064 36.625 45 33.1 45 29c0-2.252-1.488-4.33-4-6.001M24 39C12.954 39 3 34.523 3 29c0-2.252 1.488-4.33 4-6.001"/><path stroke-miterlimit="10" d="M24 20a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}