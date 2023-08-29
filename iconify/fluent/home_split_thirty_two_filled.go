package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSplitThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.108 2.776a3.25 3.25 0 0 0-4.216 0l-9.75 8.31A3.25 3.25 0 0 0 3 13.559V26.5A2.5 2.5 0 0 0 5.5 29h21a2.5 2.5 0 0 0 2.5-2.5V13.56a3.25 3.25 0 0 0-1.142-2.473l-9.75-8.31ZM17 8v2a1 1 0 1 1-2 0V8a1 1 0 1 1 2 0Zm-1 6a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1Zm1 8v2a1 1 0 1 1-2 0v-2a1 1 0 1 1 2 0Z"/>`),
		g.Group(children),
	)
}