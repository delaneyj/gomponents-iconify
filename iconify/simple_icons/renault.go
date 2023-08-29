package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Renault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.463 11.99l-4.097-7.692l-.924 1.707l3.213 5.985l-5.483 10.283L4.69 11.99L11.096 0H9.27L2.882 11.99L9.269 24h1.807zm3.655 0L14.711 0h-1.807L6.517 11.99l4.117 7.712l.904-1.707l-3.193-6.005l5.463-10.263L19.29 11.99L12.904 24h1.807Z"/>`),
		g.Group(children),
	)
}