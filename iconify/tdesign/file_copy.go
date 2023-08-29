package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 1h9.414L21 6.586V19H6V1Zm2 2v14h11V9h-6V3H8Zm7 .414V7h3.586L15 3.414ZM4 5v16h11v2H2V5h2Z"/>`),
		g.Group(children),
	)
}