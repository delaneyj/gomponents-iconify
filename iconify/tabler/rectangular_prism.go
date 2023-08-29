package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangularPrism(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 14.008V8.99a1.98 1.98 0 0 0-1-1.717l-4-2.008a2.016 2.016 0 0 0-2 0L4 10.273c-.619.355-1 1.01-1 1.718v5.018c0 .709.381 1.363 1 1.717l4 2.008a2.016 2.016 0 0 0 2 0l10-5.008c.619-.355 1-1.01 1-1.718zM9 21v-7.5m0 0L20.5 8m-17 3L9 13.5"/>`),
		g.Group(children),
	)
}