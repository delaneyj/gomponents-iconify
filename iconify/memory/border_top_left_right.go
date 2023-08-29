package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 18v2h-2v-2h2m4 0v2h-2v-2h2m-8 0v2H6v-2h2m-4 2H2V2h18v18h-2V4H4v16Z"/>`),
		g.Group(children),
	)
}