package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m960 0l960 1920H0L960 0zm64 1664v-128H896v128h128zm-128-256h128V640H896v768z"/>`),
		g.Group(children),
	)
}