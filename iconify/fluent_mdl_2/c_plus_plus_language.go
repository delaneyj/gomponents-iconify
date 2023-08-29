package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CPlusPlusLanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1408h-512v512h-128v-512H768v-128h512V768h128v512h512v128zM640 640v512H512V640H0V512h512V0h128v512h512v128H640z"/>`),
		g.Group(children),
	)
}