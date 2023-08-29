package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3V1h2v2h13.008c.548 0 .992.445.992.993v16.014a1 1 0 0 1-.992.993H2.992A.993.993 0 0 1 2 20.007V3.993A1 1 0 0 1 2.992 3H6ZM4 5v14h16V5H4Zm5 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm5-6h4v2h-4V9Zm0 4h4v2h-4v-2Z"/>`),
		g.Group(children),
	)
}