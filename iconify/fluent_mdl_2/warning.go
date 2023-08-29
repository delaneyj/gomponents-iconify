package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1920H0L960 0l960 1920zM207 1792h1506L960 286L207 1792zm817-1024v640H896V768h128zm-128 768h128v128H896v-128z"/>`),
		g.Group(children),
	)
}