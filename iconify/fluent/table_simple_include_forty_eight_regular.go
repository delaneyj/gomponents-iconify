package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleIncludeFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.25 6h25.5C39.65 6 42 8.35 42 11.25v11.417a6.241 6.241 0 0 0-2.5-1.419V11.25a2.75 2.75 0 0 0-2.75-2.75H25.5v12.748a6.241 6.241 0 0 0-2.5 1.42V8.5H11.25a2.75 2.75 0 0 0-2.75 2.75v11.5h14.413a6.253 6.253 0 0 0-1.586 2.5H8.5v11.5a2.75 2.75 0 0 0 2.75 2.75h9.998a6.241 6.241 0 0 0 1.42 2.5H11.25A5.25 5.25 0 0 1 6 36.75v-25.5C6 8.35 8.35 6 11.25 6ZM23 27.25A4.25 4.25 0 0 1 27.25 23h10.5A4.25 4.25 0 0 1 42 27.25v10.5A4.25 4.25 0 0 1 37.75 42h-10.5A4.25 4.25 0 0 1 23 37.75v-10.5Z"/>`),
		g.Group(children),
	)
}