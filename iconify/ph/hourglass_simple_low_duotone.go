package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassSimpleLowDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 216H56a8 8 0 0 1-5.66-13.66L77.87 176h100.26l27.51 26.34A8 8 0 0 1 200 216Z" opacity=".2"/><path d="M211.18 196.56L139.57 128l71.61-68.56a1.59 1.59 0 0 1 .13-.13A16 16 0 0 0 200 32H56a16 16 0 0 0-11.31 27.31a1.59 1.59 0 0 1 .13.13L116.43 128l-71.61 68.56a1.59 1.59 0 0 1-.13.13A16 16 0 0 0 56 224h144a16 16 0 0 0 11.32-27.31a1.59 1.59 0 0 1-.14-.13ZM56 48Zm102.21 120H97.79L128 139.08ZM200 48l-72 68.92L56 48ZM56 208l25.06-24h93.84l25.1 24Z"/></g>`),
		g.Group(children),
	)
}