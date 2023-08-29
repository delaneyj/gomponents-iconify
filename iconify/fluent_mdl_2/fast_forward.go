package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1152 379l806 645l-806 650V379zm128 266v762l474-383l-474-379zM256 1674V379l806 645l-806 650zM384 645v762l474-383l-474-379z"/>`),
		g.Group(children),
	)
}