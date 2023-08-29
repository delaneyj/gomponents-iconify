package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18h10v2H2v-2Zm0-7h14v2H2v-2Zm0-7h20v2H2V4Zm17 11.17V9h5v2h-3v7a3 3 0 1 1-2-2.83ZM18 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}