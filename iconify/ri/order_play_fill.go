package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrderPlayFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4V2.067a.5.5 0 0 1 .82-.384l4.12 3.433a.5.5 0 0 1-.321.884H2V4h15ZM2 18h20v2H2v-2Zm0-7h20v2H2v-2Z"/>`),
		g.Group(children),
	)
}