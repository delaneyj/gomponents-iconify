package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserToUserTransmission(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M14 18a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm20 0a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path stroke-linecap="round" d="M4 44v-9c0-5.523 3.77-10 8.421-10h5.053C21.559 25 24 29.027 24 29.027"/><path stroke-linecap="round" d="M44 44v-9c0-5.523-3.815-10-8.521-10h-5.113c-3.961 0-6.374 4.027-6.366 4.027M11 40h27"/><path stroke-linecap="round" d="m34.295 36.258l1.24 1.247L38.015 40l-2.48 2.561l-1.24 1.281m-19.965-7.61l-1.26 1.254l-2.518 2.507l2.518 2.548l1.26 1.275"/></g>`),
		g.Group(children),
	)
}