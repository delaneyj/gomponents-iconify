package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavTwoDMapView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1344 423l704-234v1414l-704 235l-640-213L0 1859V445l704-235l640 213zm-576-57v1145l512 171V537L768 366zM128 537v1145l512-171V366L128 537zm1792 974V366l-512 171v1145l512-171z"/>`),
		g.Group(children),
	)
}