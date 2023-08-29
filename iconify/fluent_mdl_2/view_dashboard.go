package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v2048H0V0h2048zm-128 128h-640v512h640V128zm-640 1152h640V768h-640v512zM128 128v1152h1024V128H128zm0 1792h640v-512H128v512zm1792 0v-512H896v512h1024z"/>`),
		g.Group(children),
	)
}