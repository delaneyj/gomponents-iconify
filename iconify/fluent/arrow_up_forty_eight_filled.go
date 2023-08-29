package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 44.25a1.5 1.5 0 0 1-1.5-1.5V10.915L10.32 23.302a1.5 1.5 0 1 1-2.14-2.104L22.928 6.201a.61.61 0 0 1 .02-.02a1.5 1.5 0 0 1 2.145.042L39.82 21.198a1.5 1.5 0 1 1-2.139 2.104L25.5 10.915V42.75a1.5 1.5 0 0 1-1.5 1.5Z"/>`),
		g.Group(children),
	)
}