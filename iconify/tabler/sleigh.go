package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19h15a4 4 0 0 0 4-4m-6 0H7a4 4 0 0 1-4-4V5l1.243 1.243A6 6 0 0 0 8.485 8H12v2a2 2 0 0 0 2 2h.5a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 1 3 0V12a3 3 0 0 1-3 3zm-1 0v4m-8-4v4"/>`),
		g.Group(children),
	)
}