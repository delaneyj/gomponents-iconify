package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitCellsHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16Zm-9 2H5v14h6v-4h2v4h6V5h-6v4h-2V5Zm4 4l3 3l-3 3v-2H9v2l-3-3l3-3v2h6V9Z"/>`),
		g.Group(children),
	)
}