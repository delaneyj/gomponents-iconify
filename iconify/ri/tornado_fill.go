package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TornadoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h20v2H2V3Zm2 4h16v2H4V7Zm4 4h14v2H8v-2Zm2 4h8v2h-8v-2Zm-2 4h6v2H8v-2Z"/>`),
		g.Group(children),
	)
}