package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 28V6h-8v22h-4V14H8v14H4V2H2v26a2 2 0 0 0 2 2h26v-2ZM22 8h4v10h-4Zm-12 8h4v6h-4Z"/>`),
		g.Group(children),
	)
}