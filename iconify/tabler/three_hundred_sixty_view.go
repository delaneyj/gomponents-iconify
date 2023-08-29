package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHundredSixtyView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-3M3 5h2.5A1.5 1.5 0 0 1 7 6.5v1A1.5 1.5 0 0 1 5.5 9H4h1.5A1.5 1.5 0 0 1 7 10.5v1A1.5 1.5 0 0 1 5.5 13H3m14-6v4a2 2 0 1 0 4 0V7a2 2 0 1 0-4 0zM3 16c0 1.657 4.03 3 9 3s9-1.343 9-3"/>`),
		g.Group(children),
	)
}