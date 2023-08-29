package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragDropLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 13l6.964 4.062l-2.973.85l2.125 3.681l-1.732 1l-2.125-3.68l-2.223 2.15L16 13Zm-2-7h2v2h5a1 1 0 0 1 1 1v4h-2v-3H10v10h4v2H9a1 1 0 0 1-1-1v-5H6v-2h2V9a1 1 0 0 1 1-1h5V6ZM4 14v2H2v-2h2Zm0-4v2H2v-2h2Zm0-4v2H2V6h2Zm0-4v2H2V2h2Zm4 0v2H6V2h2Zm4 0v2h-2V2h2Zm4 0v2h-2V2h2Z"/>`),
		g.Group(children),
	)
}