package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceCursorHandHandSelectCursorFinger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 6.13a2 2 0 0 1 1.59 2.24l-.61 4.27a1 1 0 0 1-1 .86H4a1 1 0 0 1-.93-.63L2 10.21a2 2 0 0 1 1-2.53L4.35 7V2a1.5 1.5 0 0 1 3 0v3.5Z"/>`),
		g.Group(children),
	)
}