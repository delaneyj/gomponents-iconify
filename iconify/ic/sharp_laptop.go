package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 18l2-2V4H2v12l2 2H0v2h24v-2h-4zM4 6h16v10H4V6z"/>`),
		g.Group(children),
	)
}