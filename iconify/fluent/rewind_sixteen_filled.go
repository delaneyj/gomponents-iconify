package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 4.465a1 1 0 0 0-1.576-.818L1.412 7.183a1 1 0 0 0 0 1.634l5.012 3.536A1 1 0 0 0 8 11.536V9.232l4.424 3.12A1 1 0 0 0 14 11.537V4.465a1 1 0 0 0-1.576-.818L8 6.768V4.465Z"/>`),
		g.Group(children),
	)
}