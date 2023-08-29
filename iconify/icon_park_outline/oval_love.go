package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvalLove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M29.516 17.5c5.523 9.566 4.96 20.23-1.258 23.82S12.523 40.067 7 30.5S2.04 10.27 8.258 6.68S23.994 7.934 29.516 17.5Z"/><path d="M18.258 17.5c-5.523 9.566-4.96 20.23 1.258 23.82s15.736-1.254 21.259-10.82s4.96-20.23-1.259-23.82C33.3 3.09 23.781 7.934 18.258 17.5Z"/><path d="M23.753 10.344c2.145 1.908 4.13 4.325 5.764 7.156c5.522 9.566 4.959 20.23-1.259 23.82c-1.298.75-2.74 1.132-4.26 1.18"/></g>`),
		g.Group(children),
	)
}