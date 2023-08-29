package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28 4.063l-1.625 1.25l-4.625 3.625L16.156 8l-.375-.063l-.342.22l-5.688 3.78l-4.563-.906L4 10.78V28h24V4.062zm-2 4.093v5.375l-4.22 3.345l-5.468-1.813l-.468-.156l-.406.25l-5.563 3.72L6 17.31v-4.09l3.813.75l.406.092l.342-.218l5.657-3.78l5.624.936l.437.063l.345-.282L26 8.157zm0 7.938V26H6v-6.5l3.625 1.438l.5.187l.438-.28l5.624-3.75l5.5 1.843l.5.187l.438-.344L26 16.095z"/>`),
		g.Group(children),
	)
}