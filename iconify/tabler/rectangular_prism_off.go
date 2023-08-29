package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangularPrismOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.18 8.18L4 10.273c-.619.355-1 1.01-1 1.718v5.018c0 .709.381 1.363 1 1.717l4 2.008a2.016 2.016 0 0 0 2 0l7.146-3.578m2.67-1.337l.184-.093c.619-.355 1-1.01 1-1.718V8.99a1.98 1.98 0 0 0-1-1.717l-4-2.008a2.016 2.016 0 0 0-2 0L10.854 6.84M9 21v-7.5m0 0l3.048-1.458m2.71-1.296L20.5 8m-17 3L9 13.5M3 3l18 18"/>`),
		g.Group(children),
	)
}