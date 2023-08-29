package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Insert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v896H640V0h1408zm-128 128H768v640h1152V128zM640 1152h1408v896H640v-896zm128 768h1152v-640H768v640zM109 659l366 365l-366 365l-90-90l274-275L19 749l90-90z"/>`),
		g.Group(children),
	)
}