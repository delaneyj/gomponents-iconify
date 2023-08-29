package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.044 16.04A9 9 0 0 0 7.962 3.955M5.629 5.643A9 9 0 0 0 12 21c2.491 0 4.73-1 6.36-2.631"/><path d="M11.298 11.288a1 1 0 1 0 1.402 1.427M7 12c0-1.38.559-2.629 1.462-3.534m2.607-1.38C11.371 7.03 11.682 7 12 7m0 10a4.985 4.985 0 0 0 3.551-1.48m1.362-2.587c.057-.302.087-.614.087-.933M3 3l18 18"/></g>`),
		g.Group(children),
	)
}