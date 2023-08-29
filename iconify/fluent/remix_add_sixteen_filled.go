package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemixAddSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 1.751a.75.75 0 0 1 .75-.75l7.911.198A7.003 7.003 0 0 1 15 8.001v.002a6.972 6.972 0 0 1-1.639 4.498h-2.2A5.5 5.5 0 0 0 8 2.501H1.751a.75.75 0 0 1-.75-.75Zm0 6.25a7 7 0 0 0 7 7h6.25a.75.75 0 0 0 0-1.5H8a5.503 5.503 0 0 1-3.164-10H2.637A6.974 6.974 0 0 0 1 8.001ZM8 5a.75.75 0 0 1 .75.75v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5A.75.75 0 0 1 8 5Z"/>`),
		g.Group(children),
	)
}