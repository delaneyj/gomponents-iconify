package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 7.25C7.5 6.56 8.06 6 8.75 6h32.5a1.25 1.25 0 0 1 1.007 1.99L33.801 19.5l8.456 11.51A1.25 1.25 0 0 1 41.25 33H10v9.75a1.25 1.25 0 1 1-2.5 0V7.25ZM10 30.5h28.78l-7.537-10.26a1.25 1.25 0 0 1 0-1.48L38.78 8.5H10v22Z"/>`),
		g.Group(children),
	)
}