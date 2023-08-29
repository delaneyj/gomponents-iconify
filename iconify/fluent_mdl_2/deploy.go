package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deploy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m960 870l301 301l-90 90l-147-146v677H896v-677l-147 146l-90-90l301-301zM768 128h512v512H768V128zm128 384h256V256H896v256zM128 128h512v512H128V128zm1792 0v512h-512V128h512z"/>`),
		g.Group(children),
	)
}