package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v640h-128V347L347 1792h421v128H128v-640h128v421L1701 256h-421V128h640z"/>`),
		g.Group(children),
	)
}