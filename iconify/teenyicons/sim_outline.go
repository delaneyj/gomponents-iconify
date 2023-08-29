package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 5.5h3.5V12M4 8.5h2m-2 3h2m3-6h2m-2 3h2m-2 3h2m1.5 3h-10a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1h7l4 4v9a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}