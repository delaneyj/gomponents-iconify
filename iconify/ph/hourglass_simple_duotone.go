package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M205.64 53.66L128 128L50.36 53.66A8 8 0 0 1 56 40h144a8 8 0 0 1 5.64 13.66ZM128 128l-77.64 74.34A8 8 0 0 0 56 216h144a8 8 0 0 0 5.66-13.66Z" opacity=".2"/><path d="M211.18 196.56L139.57 128l71.61-68.56a1.59 1.59 0 0 1 .13-.13A16 16 0 0 0 200 32H56a16 16 0 0 0-11.31 27.31a1.59 1.59 0 0 1 .13.13L116.43 128l-71.61 68.56a1.59 1.59 0 0 1-.13.13A16 16 0 0 0 56 224h144a16 16 0 0 0 11.32-27.31a1.59 1.59 0 0 1-.14-.13ZM56 48Zm144 0l-72 68.92L56 48ZM56 208l72-68.92L200 208Z"/></g>`),
		g.Group(children),
	)
}