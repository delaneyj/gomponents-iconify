package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.585 15.5H5.415A1.65 1.65 0 0 1 4 13a10.526 10.526 0 0 0 1.5-5.415V6.5a4 4 0 0 1 4-4h2a4 4 0 0 1 4 4v1.085c0 1.907.518 3.78 1.5 5.415a1.65 1.65 0 0 1-1.415 2.5zM13 17c-.667 1-1.5 1.5-2.5 1.5S8.667 18 8 17"/>`),
		g.Group(children),
	)
}