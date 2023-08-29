package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 6a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Zm2 5a4.481 4.481 0 0 0 2.809-.984l3.837 3.838a.5.5 0 0 0 .708-.708L10.016 9.31A4.5 4.5 0 1 0 6.5 11Zm0-8a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7Z"/>`),
		g.Group(children),
	)
}