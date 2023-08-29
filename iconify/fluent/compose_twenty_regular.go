package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComposeTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.854 2.854a.5.5 0 0 0-.707-.708l-8 8L9 11l.854-.146l8-8ZM5.5 3A2.5 2.5 0 0 0 3 5.5v9A2.5 2.5 0 0 0 5.5 17h9a2.5 2.5 0 0 0 2.5-2.5v-6a.5.5 0 0 0-1 0v6a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 4 14.5v-9A1.5 1.5 0 0 1 5.5 4h6.005a.5.5 0 0 0 0-1H5.5Z"/>`),
		g.Group(children),
	)
}