package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickCursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.974 16.595L4.5 21.828l17.336-10.301"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.836 11.528c2.654-1.138 5.814 8.59 5.814 8.59c.642-1.669 2.504-1.797 3.766-1.455a2.61 2.61 0 0 1 3.96-.686c.575-1.881 2.8-2.138 3.461-.85l4.107 9.164c3.035 6.922-7.202 11.752-10.394 9.86l-12.206-7.299c-2.269-1.5-2.273-4.572 1.6-3.41l5.918 3.46s-9.156-15.894-6.026-17.37Z"/>`),
		g.Group(children),
	)
}