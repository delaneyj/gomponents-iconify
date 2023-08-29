package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoupOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19h16m-4-8h6c0 1.691-.525 3.26-1.42 4.552m-2.034 2.032A7.963 7.963 0 0 1 13 19h-2a8 8 0 0 1-8-8h8m1-6v3m3-3v3M3 3l18 18"/>`),
		g.Group(children),
	)
}