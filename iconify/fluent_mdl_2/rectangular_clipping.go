package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangularClipping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1152h-128V384H128v1280h1024v128H0V256h2048zm-256 1408h256v128h-256v256h-128v-256h-256v-128h256v-256h128v256z"/>`),
		g.Group(children),
	)
}