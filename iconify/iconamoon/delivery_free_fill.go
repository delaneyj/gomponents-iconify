package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeliveryFreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4a2 2 0 0 1 2-2a3.99 3.99 0 0 1 3 1.354A3.99 3.99 0 0 1 10.5 2a2 2 0 0 1 2 2c0 .729-.195 1.412-.535 2H13a1 1 0 0 1 1 1v1h4a5 5 0 0 1 5 5v4a3.001 3.001 0 0 1-2.129 2.872a3 3 0 0 1-5.7.128H8.83a3 3 0 0 1-5.7-.128A3.001 3.001 0 0 1 1 17V7a1 1 0 0 1 1-1h1.035A3.982 3.982 0 0 1 2.5 4Zm2 0a2 2 0 0 1 2 2a2 2 0 0 1-2-2Zm4 2a2 2 0 0 0 2-2a2 2 0 0 0-2 2Zm6.671 12H14v-8h4a3 3 0 0 1 3 3v4a.997.997 0 0 1-.293.707a3 3 0 0 0-5.536.293Zm-9.878.293a1 1 0 1 1 1.414 1.414a1 1 0 0 1-1.414-1.414ZM17 19a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}