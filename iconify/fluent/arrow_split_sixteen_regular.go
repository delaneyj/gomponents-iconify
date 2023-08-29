package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSplitSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a.5.5 0 0 1 .5.5V6h2A1.5 1.5 0 0 1 12 7.5v4.793l1.146-1.147a.5.5 0 0 1 .708.708l-2 2a.5.5 0 0 1-.708 0l-2-2a.5.5 0 0 1 .708-.708L11 12.293V7.5a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.5.5v4.793l1.146-1.147a.5.5 0 0 1 .708.708l-2 2a.5.5 0 0 1-.708 0l-2-2a.5.5 0 0 1 .708-.708L4 12.293V7.5A1.5 1.5 0 0 1 5.5 6h2V2.5A.5.5 0 0 1 8 2Z"/>`),
		g.Group(children),
	)
}