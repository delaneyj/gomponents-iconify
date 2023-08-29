package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.25 4A2.25 2.25 0 0 0 2 6.25V8h18.5v11.25a.75.75 0 0 0 1.5 0v-13A2.25 2.25 0 0 0 19.75 4H4.25ZM12 9.5H2v7.75A2.75 2.75 0 0 0 4.75 20h4.5A2.75 2.75 0 0 0 12 17.25V9.5Zm-7 3.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}