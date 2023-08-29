package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductHuntFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 22c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Zm1.334-10H10.5V9h2.834a1.5 1.5 0 0 1 0 3Zm0-5H8.5v10h2v-3h2.834a3.5 3.5 0 1 0 0-7Z"/>`),
		g.Group(children),
	)
}