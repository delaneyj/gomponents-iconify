package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeaderFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 1024V384h128v1408H768v-640H128v640H0V384h128v640h640zm1152 256h128v128h-128v384h-128v-384h-640v-85l671-939h97v896zm-128 0V648l-452 632h452z"/>`),
		g.Group(children),
	)
}