package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abacus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v18m14 0V3M5 7h14M5 15h14M8 13v4m3-4v4m5-4v4M14 5v4m-3-4v4M8 5v4M3 21h18"/>`),
		g.Group(children),
	)
}