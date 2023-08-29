package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarVerticalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 3.5a1.5 1.5 0 1 1 3 0v9a1.5 1.5 0 0 1-3 0v-9ZM3.5 3a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 1 0v-9a.5.5 0 0 0-.5-.5ZM6 6.5a1.5 1.5 0 1 1 3 0v6a1.5 1.5 0 0 1-3 0v-6ZM7.5 6a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 1 0v-6a.5.5 0 0 0-.5-.5Zm4 2A1.5 1.5 0 0 0 10 9.5v3a1.5 1.5 0 0 0 3 0v-3A1.5 1.5 0 0 0 11.5 8ZM11 9.5a.5.5 0 0 1 1 0v3a.5.5 0 0 1-1 0v-3Z"/>`),
		g.Group(children),
	)
}