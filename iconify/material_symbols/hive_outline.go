package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HiveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.75 22l-1.7-3h-3.3L2.5 15l1.7-3l-1.7-3l2.25-4h3.3l1.7-3h4.5l1.7 3h3.3l2.25 4l-1.7 3l1.7 3l-2.25 4h-3.3l-1.7 3h-4.5Zm6.2-11h2.15l1.1-2l-1.1-2h-2.15l-1.125 2l1.125 2Zm-5 3h2.1l1.125-2l-1.125-2h-2.1l-1.125 2l1.125 2Zm0-6h2.1l1.15-2.025L13.075 4h-2.15L9.8 5.975L10.95 8Zm-5.025 3H8.05l1.125-2L8.05 7H5.925L4.8 9l1.125 2Zm0 6H8.05l1.125-2l-1.125-2H5.9l-1.1 2l1.125 2Zm5 3h2.15l1.125-1.975L13.05 16h-2.1L9.8 18.025L10.925 20Zm5.025-3h2.125l1.125-2l-1.125-2H15.95l-1.125 2l1.125 2Z"/>`),
		g.Group(children),
	)
}