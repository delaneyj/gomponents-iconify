package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCellMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2H2v2h2V2Zm2 0v2h2V2h2v2h2V2H6Zm6 4H6v6h6V6Zm-8 6v-2H2v2h2ZM2 8h2V6H2v2Zm0-8h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}