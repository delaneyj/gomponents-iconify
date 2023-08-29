package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cleanfox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 32.208c-2.491-.766-4.552-1.245-4.552-1.245c0-2.923.144-6.468-1.197-9.103c1.293-3.45-.689-12.991-3.546-13.607c-2.443-.527-6.132 5.797-6.132 5.797a8.058 8.058 0 0 0-2.156-.614V11.87l-1.556 1.419l.621-3.168l-3.182 3.365a7.76 7.76 0 0 0-1.873.564s-3.689-6.324-6.132-5.797c-2.857.616-4.84 10.157-3.546 13.607c-1.341 2.635-1.197 6.18-1.197 9.103c0 0-2.06.48-4.552 1.246a10.73 10.73 0 0 0 4.648 1.437s-.815.288-1.677.575C9.195 35.898 16.526 39.78 24 39.78s14.804-3.881 16.529-5.558c-.862-.288-1.677-.575-1.677-.575a10.73 10.73 0 0 0 4.647-1.438Z"/>`),
		g.Group(children),
	)
}