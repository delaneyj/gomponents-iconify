package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 4.465a1 1 0 0 0-1.576-.818L1.412 7.183a1 1 0 0 0 0 1.634l5.012 3.536A1 1 0 0 0 8 11.536V9.232l4.424 3.12A1 1 0 0 0 14 11.537V4.465a1 1 0 0 0-1.576-.818L8 6.768V4.465Zm0 3.543v-.016l5-3.527v7.07L8 8.009Zm-1 3.528L1.988 8L7 4.465v7.07Z"/>`),
		g.Group(children),
	)
}