package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scribble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15c2 3 4 4 7 4s7-3 7-7s-3-7-6-7s-5 1.5-5 4s2 5 6 5s8.408-2.453 10-5"/>`),
		g.Group(children),
	)
}