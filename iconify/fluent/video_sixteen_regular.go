package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 5a1.5 1.5 0 0 1 1.5-1.5h6A1.5 1.5 0 0 1 11 5v6a1.5 1.5 0 0 1-1.5 1.5h-6A1.5 1.5 0 0 1 2 11V5Zm1.5-2.5A2.5 2.5 0 0 0 1 5v6a2.5 2.5 0 0 0 2.5 2.5h6A2.5 2.5 0 0 0 12 11v-.5l1.8 1.35a.75.75 0 0 0 1.2-.6V4.755a.75.75 0 0 0-1.2-.6L12 5.5V5a2.5 2.5 0 0 0-2.5-2.5h-6ZM12 6.75l2-1.496v5.496l-2-1.5v-2.5Z"/>`),
		g.Group(children),
	)
}