package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acadboost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10"><path stroke-width=".866" d="M33.5 27.2H14.8m19.7 0h0c4.2 0 7.6 3.4 7.6 7.6s-3.4 7.6-7.6 7.6H23.9"/><path stroke-width=".866" d="M23.9 42.5V12h9.6c4.2 0 7.6 3.4 7.6 7.6s-3.4 7.6-7.6 7.6v0"/><path d="M5.9 42.4L27.7 5.5m-4.5 1.1l4.5-1.1l1.2 4.5"/></g>`),
		g.Group(children),
	)
}