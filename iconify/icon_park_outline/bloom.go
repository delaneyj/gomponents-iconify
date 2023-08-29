package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bloom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 32l12 10m24-10L30 42m-6-10v12m-7-33l7-7l7 7l7-1s1 4.239 1 7c0 10-8.5 15-15 15S9 27 9 17c0-2.761 1-7 1-7l7 1Z"/>`),
		g.Group(children),
	)
}