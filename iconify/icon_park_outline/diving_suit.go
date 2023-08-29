package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivingSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M17 24h-2.91C8.52 24 4 19.52 4 14S8.52 4 14.09 4L20 9l5.91-5C31.48 4 36 8.48 36 14s-4.52 10-10.09 10H23m3 12H14v8h12v-8Z"/><path d="M26 40h10c4.42 0 8-3.58 8-8V14h-8M21.22 30h-2.44c-1.46 0-2.57-1.31-2.33-2.74l1.22-7.29A2.354 2.354 0 0 1 20 18c1.15 0 2.14.83 2.33 1.97l1.22 7.29c.23 1.43-.88 2.74-2.33 2.74Z"/></g>`),
		g.Group(children),
	)
}