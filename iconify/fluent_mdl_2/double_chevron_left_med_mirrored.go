package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronLeftMedMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m903 146l879 878l-879 878l121 121l999-999l-999-999l-121 121zm-853 0l878 878l-878 878l121 121l999-999L171 25L50 146z"/>`),
		g.Group(children),
	)
}