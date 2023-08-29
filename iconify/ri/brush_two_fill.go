package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.536 15.947l2.121-2.122l-3.182-3.182l3.536-3.535l-2.122-2.122l-3.535 3.536l-3.182-3.182L8.05 7.461l8.486 8.486ZM13.354 5.694l2.828-2.829a1 1 0 0 1 1.414 0l3.536 3.536a1 1 0 0 1 0 1.414l-2.829 2.828l2.475 2.475a1 1 0 0 1 0 1.414L13 22.311a1 1 0 0 1-1.414 0l-9.9-9.9a1 1 0 0 1 0-1.414l7.779-7.778a1 1 0 0 1 1.414 0l2.475 2.475Z"/>`),
		g.Group(children),
	)
}