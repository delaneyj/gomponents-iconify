package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReviewSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m448 768l-320 320V768H0V128h1664v640H448zm-64 256h1664v640h-128v320l-320-320H384v-640z"/>`),
		g.Group(children),
	)
}