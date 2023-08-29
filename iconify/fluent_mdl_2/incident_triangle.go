package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncidentTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 2048H0L960 128l960 1920zm-896-384H896v128h128v-128zm0-128V896H896v640h128z"/>`),
		g.Group(children),
	)
}