package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAdditionSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 0v1920H0V0h1920zm-384 896h-512V384H896v512H384v128h512v512h128v-512h512V896z"/>`),
		g.Group(children),
	)
}