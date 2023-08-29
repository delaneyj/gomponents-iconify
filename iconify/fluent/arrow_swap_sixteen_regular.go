package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSwapSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.354 1.646a.5.5 0 0 0-.708.708L11.293 4H3.5a.5.5 0 0 0 0 1h7.793L9.646 6.646a.5.5 0 1 0 .708.708l2.5-2.5a.5.5 0 0 0 0-.708l-2.5-2.5Zm-4 7.708a.5.5 0 1 0-.708-.708l-2.5 2.5a.5.5 0 0 0 0 .708l2.5 2.5a.5.5 0 0 0 .708-.708L4.707 12H12.5a.5.5 0 0 0 0-1H4.707l1.647-1.646Z"/>`),
		g.Group(children),
	)
}