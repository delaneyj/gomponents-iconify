package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeCellsHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16Zm-9 2H5v5.999h2V9l3 3l-3 3v-2H5v6h6v-2h2v2h6v-6h-2v2l-3-3l3-3v1.999h2V5h-6v2h-2V5Zm2 8v2h-2v-2h2Zm0-4v2h-2V9h2Z"/>`),
		g.Group(children),
	)
}