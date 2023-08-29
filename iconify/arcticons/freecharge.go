package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freecharge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.383 43l-9.14-15.81l-18.631-4.986L41.388 5l-1.345 8.218l-12.481 6.128l2.03 3.513l9.41-4.86l-1.3 8.088l-4.583 2.223l2.86 4.947Z"/>`),
		g.Group(children),
	)
}