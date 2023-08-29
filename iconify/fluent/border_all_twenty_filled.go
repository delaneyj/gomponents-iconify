package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAllTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 4.5A1.75 1.75 0 0 0 4.5 6.25v3h4.75V4.5h-3Zm4.5 0v4.75h4.75v-3a1.75 1.75 0 0 0-1.75-1.75h-3Zm4.75 6.25h-4.75v4.75h3a1.75 1.75 0 0 0 1.75-1.75v-3ZM9.25 15.5v-4.75H4.5v3c0 .966.784 1.75 1.75 1.75h3ZM3 6.25A3.25 3.25 0 0 1 6.25 3h7.5A3.25 3.25 0 0 1 17 6.25v7.5A3.25 3.25 0 0 1 13.75 17h-7.5A3.25 3.25 0 0 1 3 13.75v-7.5Z"/>`),
		g.Group(children),
	)
}