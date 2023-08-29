package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 21h8m-4-6v6M7 3h10l1 7a4.511 4.511 0 0 1-1.053 2.94m-2.386 1.625A7.48 7.48 0 0 1 12 15c-3.314 0-6-1.988-6-5l.5-3.495M3 3l18 18"/>`),
		g.Group(children),
	)
}