package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowWrapTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 4a3 3 0 1 1 0 6H6a2 2 0 1 0 0 4h10.293l-.647-.646a.5.5 0 0 1 .708-.708l1.5 1.5a.5.5 0 0 1 0 .708l-1.5 1.5a.5.5 0 0 1-.708-.708l.647-.646H6a3 3 0 1 1 0-6h8a2 2 0 1 0 0-4H3.5a.5.5 0 0 1 0-1H14Z"/>`),
		g.Group(children),
	)
}