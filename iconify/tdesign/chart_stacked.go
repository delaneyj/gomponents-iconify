package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h2V9h6v11h2V5h6v15h2v2H2V2h2Zm14 18v-3.5h-2V20h2Zm-2-5.5h2V7h-2v7.5ZM10 20v-3.5H8V20h2Zm-2-5.5h2V11H8v3.5Z"/>`),
		g.Group(children),
	)
}