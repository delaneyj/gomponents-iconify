package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottomLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 4V2h2v2h-2M6 4V2h2v2H6m8 0V2h2v2h-2m4-2h2v18H2V2h2v16h14V2Z"/>`),
		g.Group(children),
	)
}