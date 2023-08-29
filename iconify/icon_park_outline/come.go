package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Come(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 40C10.128 35.564 7 26 4 19.5c-1.35-2.924 3.526-3.69 6.5-2.5c2.5 1 5.5 3 5.5 3v-8.5a3.5 3.5 0 1 1 7 0v-2a3.5 3.5 0 1 1 7 0v4a3.5 3.5 0 1 1 7 0v3a3.5 3.5 0 0 1 7-.002V29c0 3.5-2 8-7 11c-4.793 2.876-12 3-18 0Z"/>`),
		g.Group(children),
	)
}