package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceRun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 384v896h1024v128h-128v128h128v128H640v-128h256v-128H0V256h1920v1093l-128-76V384H128zm1152 768l747 448l-747 448v-896zm128 226v444l370-222l-370-222z"/>`),
		g.Group(children),
	)
}