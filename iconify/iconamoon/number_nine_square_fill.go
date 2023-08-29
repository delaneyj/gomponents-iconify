package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14Zm7.164 12.45a1 1 0 0 0 1.672 1.097l3.489-5.323a4 4 0 1 0-3.55 1.769l-1.611 2.458ZM14 10a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}