package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 28h-6a2.002 2.002 0 0 1-2-2v-5H6a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h5V6a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v5h5a2.002 2.002 0 0 1 2 2v6a2.003 2.003 0 0 1-2 2h-5v5a2.003 2.003 0 0 1-2 2ZM6 13v6h7v7h6v-7h7v-6h-7V6h-6v7Z"/>`),
		g.Group(children),
	)
}