package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M8 40V10c0-3.314 2.865-6 6.4-6H40v32H14.4c-5.038 0-6.4.684-6.4 4Z"/><path stroke-linecap="round" d="M12 44h28v-8H12a4 4 0 0 0 0 8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}