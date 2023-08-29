package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enquire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-miterlimit="2" d="M37 16a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm-25-4a4 4 0 1 1 0-8a4 4 0 0 1 0 8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m26 39l6-5v-6c0-3.466 2-6 5-6s5 2.534 5 6v14m-18-9l-6-5v-4c0-3.466-2-6-5-6s-5 2.534-5 6v18"/></g>`),
		g.Group(children),
	)
}