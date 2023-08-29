package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextboxAlignBottomRotateNinetyTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 5.5A2.5 2.5 0 0 0 14.5 3h-9A2.5 2.5 0 0 0 3 5.5v9A2.5 2.5 0 0 0 5.5 17h9a2.5 2.5 0 0 0 2.5-2.5v-9Zm-7 1v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0Zm-3 0v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}