package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 5h9v2h-7v10.172l3.95-3.95l1.414 1.414L11 21l-6.364-6.364l1.414-1.414l3.95 3.95V5Z"/>`),
		g.Group(children),
	)
}