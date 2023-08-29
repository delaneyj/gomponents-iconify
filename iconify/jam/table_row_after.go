package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRowAfter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6v6h1a1 1 0 0 1 0 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1a1 1 0 0 1 0-2h1V6H2Zm0-2h2V2H2v2Zm4-2v2h2V2H6Zm4 0v2h2V2h-2Zm-2 9v1a1 1 0 0 1-2 0v-1H5a1 1 0 0 1 0-2h1V8a1 1 0 1 1 2 0v1h1a1 1 0 1 1 0 2H8Z"/>`),
		g.Group(children),
	)
}