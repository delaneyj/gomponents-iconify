package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.39 1.73L1.11 3l4 4H4l-1 5v2h1v6h10v-4.11l6.84 6.84l1.27-1.27L2.39 1.73M5.64 9h1.47l3 3H5.04l.6-3M12 18H6v-4h6v4m6-3.2V14h-.8l-2-2h3.76l-.6-3H12.2l-2-2H20l1 5v2h-1v2.8l-2-2M9.2 6l-2-2H20v2H9.2Z"/>`),
		g.Group(children),
	)
}