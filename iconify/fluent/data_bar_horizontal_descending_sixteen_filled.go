package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarHorizontalDescendingSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 5a2 2 0 1 0 0-4H3a2 2 0 1 0 0 4h10ZM7 15a2 2 0 1 0 0-4H3a2 2 0 1 0 0 4h4Zm4-7a2 2 0 0 1-2 2H3a2 2 0 1 1 0-4h6a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}