package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewAllTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 0v896H0V0h896zM768 768V128H128v640h640zM0 1920v-896h1920v896H0zm128-768v640h1664v-640H128zM1024 0h896v896h-896V0zm768 768V128h-640v640h640z"/>`),
		g.Group(children),
	)
}