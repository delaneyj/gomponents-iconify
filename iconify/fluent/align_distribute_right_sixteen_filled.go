package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignDistributeRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 1.5a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13Zm-7 0a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13ZM3.5 13A1.5 1.5 0 0 1 2 11.5v-7A1.5 1.5 0 0 1 3.5 3h1A1.5 1.5 0 0 1 6 4.5v7A1.5 1.5 0 0 1 4.5 13h-1ZM9 9.5a1.5 1.5 0 0 0 1.5 1.5h1A1.5 1.5 0 0 0 13 9.5v-3A1.5 1.5 0 0 0 11.5 5h-1A1.5 1.5 0 0 0 9 6.5v3Z"/>`),
		g.Group(children),
	)
}