package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 29a1 1 0 1 0 2 0V3a1 1 0 1 0-2 0v26Zm11.5-3h-8V6h8A3.5 3.5 0 0 1 30 9.5v13a3.5 3.5 0 0 1-3.5 3.5Zm-13-20h-8A3.5 3.5 0 0 0 2 9.5v13A3.5 3.5 0 0 0 5.5 26h8V6Z"/>`),
		g.Group(children),
	)
}