package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitrue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29 40.613l11.887-6.863v-19.5M7.113 20.024V33.75L24 43.5m11.887-32.137L24 4.5L7.113 14.25"/><circle cx="24" cy="27.239" r="6.95" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.05 14.916v12.323"/>`),
		g.Group(children),
	)
}