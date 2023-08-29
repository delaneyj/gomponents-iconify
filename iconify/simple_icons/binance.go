package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.624 13.92l2.717 2.716l-7.353 7.353l-7.352-7.352l2.717-2.717l4.636 4.66l4.635-4.66zm4.637-4.636L24 12l-2.715 2.716L18.568 12l2.693-2.716zm-9.272 0l2.716 2.692l-2.717 2.717L9.272 12l2.716-2.715zm-9.273 0L5.41 12l-2.692 2.692L0 12l2.716-2.716zM11.99.01l7.352 7.33l-2.717 2.715l-4.636-4.636l-4.635 4.66l-2.717-2.716L11.989.011z"/>`),
		g.Group(children),
	)
}