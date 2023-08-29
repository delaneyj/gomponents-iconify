package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6a2 2 0 0 0-3.562-1.25l-2.789 3.486L11.2 6.4a2 2 0 0 0-2.762.351l-4 5A1.998 1.998 0 0 0 4 13v3h16V6zm0 13H4a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2z"/>`),
		g.Group(children),
	)
}