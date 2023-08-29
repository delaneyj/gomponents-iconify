package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitCellsVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16Zm-1 2H5v5.999L9 11v2H5v6h14v-6h-4v-2l4-.001V5Zm-7 1l3 3h-2v6h2l-3 3l-3-3h2V9H9l3-3Z"/>`),
		g.Group(children),
	)
}