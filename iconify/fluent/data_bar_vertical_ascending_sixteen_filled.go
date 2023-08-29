package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarVerticalAscendingSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 3a2 2 0 1 1 4 0v10a2 2 0 1 1-4 0V3ZM1 9a2 2 0 1 1 4 0v4a2 2 0 1 1-4 0V9Zm7-4a2 2 0 0 0-2 2v6a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}