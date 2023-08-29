package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeilingLightMultipleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.76 13l2 4H9.24l2-4h5.52M15 6h-2v5h-3l-4 8h16l-4-8h-3V6m1 14c0 1.11-.89 2-2 2s-2-.89-2-2h4m-7.79-9.89L8.76 9H11V2H9v5H6l-4 8h3.76l2.45-4.89Z"/>`),
		g.Group(children),
	)
}