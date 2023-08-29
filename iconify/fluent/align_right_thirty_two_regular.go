package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M29 4a1 1 0 1 0-2 0v24a1 1 0 1 0 2 0V4Zm-7.5 1A3.5 3.5 0 0 1 25 8.5v3a3.5 3.5 0 0 1-3.5 3.5h-15A3.5 3.5 0 0 1 3 11.5v-3A3.5 3.5 0 0 1 6.5 5h15ZM23 8.5A1.5 1.5 0 0 0 21.5 7h-15A1.5 1.5 0 0 0 5 8.5v3A1.5 1.5 0 0 0 6.5 13h15a1.5 1.5 0 0 0 1.5-1.5v-3ZM21.5 17a3.5 3.5 0 0 1 3.5 3.5v3a3.5 3.5 0 0 1-3.5 3.5h-9A3.5 3.5 0 0 1 9 23.5v-3a3.5 3.5 0 0 1 3.5-3.5h9Zm1.5 3.5a1.5 1.5 0 0 0-1.5-1.5h-9a1.5 1.5 0 0 0-1.5 1.5v3a1.5 1.5 0 0 0 1.5 1.5h9a1.5 1.5 0 0 0 1.5-1.5v-3Z"/>`),
		g.Group(children),
	)
}