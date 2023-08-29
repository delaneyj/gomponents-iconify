package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4a12 12 0 1 0 12 12A12 12 0 0 0 16 4Zm10 11h-9V6.05A10 10 0 0 1 26 15ZM15.42 26A10 10 0 0 1 15 6.05v9a2 2 0 0 0 2 2h9A10 10 0 0 1 15.42 26Z"/>`),
		g.Group(children),
	)
}