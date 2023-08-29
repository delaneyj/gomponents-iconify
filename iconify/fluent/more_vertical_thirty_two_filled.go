package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 7a3 3 0 1 0 6 0a3 3 0 0 0-6 0Zm0 9a3 3 0 1 0 6 0a3 3 0 0 0-6 0Zm0 9a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z"/>`),
		g.Group(children),
	)
}