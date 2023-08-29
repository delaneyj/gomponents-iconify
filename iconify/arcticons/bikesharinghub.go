package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bikesharinghub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.327 16.096h3.142l-3.197 11.93h10.943l-8.459-8.595m-2.484 8.594L14.58 17.373l-2.794 10.429l4.05-15.113l4.052 1.086"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.186 26.097a7.243 7.243 0 0 1 .743-1.759m2.921-2.777a7.286 7.286 0 1 1-3.607 8.587m-13.658-8.565a7.287 7.287 0 1 1-3.8-1.068"/>`),
		g.Group(children),
	)
}