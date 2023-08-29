package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeliveryFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 2a1 1 0 0 0-1 1v14c0 1.354.897 2.498 2.129 2.872a3 3 0 0 0 5.7.128h6.341a3 3 0 0 0 5.7-.128A3.001 3.001 0 0 0 23 17v-4a5 5 0 0 0-5-5h-4V3a1 1 0 0 0-1-1H2Zm13.171 16H14v-8h4a3 3 0 0 1 3 3v4a.997.997 0 0 1-.293.707a3 3 0 0 0-5.536.293Zm-9.878.293a1 1 0 1 1 1.414 1.414a1 1 0 0 1-1.414-1.414ZM17 19a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}