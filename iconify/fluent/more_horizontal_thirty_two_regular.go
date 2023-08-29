package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHorizontalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 16a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm9 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm6.5 2.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}