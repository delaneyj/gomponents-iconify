package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 15h10a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H32M17 15H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h11m7-27v36M12 24h4m16 0h4"/>`),
		g.Group(children),
	)
}