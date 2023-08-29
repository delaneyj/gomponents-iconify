package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretVerticalSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 3.375L11 6H4l3.5-2.625Zm0 8.25L4 9h7l-3.5 2.625Z"/>`),
		g.Group(children),
	)
}