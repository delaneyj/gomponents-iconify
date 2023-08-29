package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemixAddTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 1.999a.75.75 0 0 0 0 1.5H10a6.503 6.503 0 0 1 3.466 12h2.344A8 8 0 0 0 10 2H2.75Zm5.598 15.83A8.003 8.003 0 0 1 4.19 4.498h2.344A6.495 6.495 0 0 0 3.5 10a6.471 6.471 0 0 0 1.391 4.02A6.49 6.49 0 0 0 10 16.498h7.25a.75.75 0 0 1 0 1.5H10a8.04 8.04 0 0 1-1.652-.17ZM10 7a.75.75 0 0 1 .75.75v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5A.75.75 0 0 1 10 7Z"/>`),
		g.Group(children),
	)
}