package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2h6v5h5v6h-5v9H9v-9H4V7h5V2Zm2 2v5H6v2h5v9h2v-9h5V9h-5V4h-2Z"/>`),
		g.Group(children),
	)
}