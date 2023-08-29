package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneAddChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7h2v10h-2zm4 6h2v4h-2z"/><path fill="currentColor" d="M19 19H5V5h9V3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2v-9h-2v9z"/><path fill="currentColor" d="M7 10h2v7H7zm12-5V3h-2v2h-2v2h2v2h2V7h2V5z"/>`),
		g.Group(children),
	)
}