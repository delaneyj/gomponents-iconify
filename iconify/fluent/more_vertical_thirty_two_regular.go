package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 9.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm0 9a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5ZM13.5 25a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Z"/>`),
		g.Group(children),
	)
}