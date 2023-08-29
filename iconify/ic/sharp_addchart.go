package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAddchart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 9h2v8h-2V9zm-2 8v-6H7v6h2zm10 2H5V5h6V3H3v18h18v-8h-2v6zm-4-6v4h2v-4h-2zm4-8V2h-2v3h-3v2h3v3h2V7h3V5h-3z"/>`),
		g.Group(children),
	)
}