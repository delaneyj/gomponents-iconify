package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SatelliteBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20.47 10.918s-1.848-.615-4.31-3.078c-2.463-2.463-3.079-4.31-3.079-4.31"/><path stroke-linecap="round" d="M20.47 10.918a5.224 5.224 0 0 1-7.388-7.388m7.388 7.388A5.224 5.224 0 0 0 18 2.145m2.47 8.773L17.646 16M13.082 3.53A5.204 5.204 0 0 1 15 2.31m-1.918 1.22L6.836 7m9.324.84L11 13M2 9.687l3-1.666M14.313 22l2.187-4M5 19l4-4"/></g>`),
		g.Group(children),
	)
}