package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeCellsVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 20a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v16Zm-2-9V5h-5.999v2H15l-3 3l-3-3h2V5H5v6h2v2H5v6h6v-2H9l3-3l3 3h-1.999v2H19v-6h-2v-2h2Zm-8 2H9v-2h2v2Zm4 0h-2v-2h2v2Z"/>`),
		g.Group(children),
	)
}