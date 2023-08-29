package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h2V9h6v11h2V5h6v15h2v2H2V2h2Zm14 18V7h-2v13h2Zm-8 0v-9H8v9h2Z"/>`),
		g.Group(children),
	)
}