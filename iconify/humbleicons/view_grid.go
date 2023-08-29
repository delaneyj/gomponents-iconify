package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 10.5V4h6v6.5H4Zm10 0V4h6v6.5h-6Zm-10 10V14h6v6.5H4Zm10 0V14h6v6.5h-6Z"/>`),
		g.Group(children),
	)
}