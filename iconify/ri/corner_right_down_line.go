package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 5H5v2h7v10.172l-3.95-3.95l-1.414 1.414L13 21l6.364-6.364l-1.414-1.414l-3.95 3.95V5Z"/>`),
		g.Group(children),
	)
}