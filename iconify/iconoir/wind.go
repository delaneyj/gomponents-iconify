package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.279 7C19.782 7 21 8.12 21 9.5S19.782 12 18.279 12H3m14.938 8c1.139 0 2.562-.5 2.562-2.5S19.077 15 17.937 15H3m7.412-11C11.842 4 13 5.12 13 6.5S11.841 9 10.412 9H3"/>`),
		g.Group(children),
	)
}