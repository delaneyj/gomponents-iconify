package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipkartAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.378 19.84h-7.86L25.75 6.842a2 2 0 0 0-3.5 0L14.481 19.84H6.622a3.12 3.12 0 0 0-3 3.93l4.38 16.308a2.84 2.84 0 0 0 2.74 2.1h26.497a2.84 2.84 0 0 0 2.76-2.1l4.41-16.308a3.12 3.12 0 0 0-3.03-3.93ZM24 11.841l4.77 8h-9.54Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.754 42.102l2.87-13.377c.168-.783.752-2.717 2.653-2.601s1.412.343 1.412.343M25.82 31.44h9.339m-9.339 0h9.339"/>`),
		g.Group(children),
	)
}