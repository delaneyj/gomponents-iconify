package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speedboat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 17h13.4a3 3 0 0 0 2.5-1.34L22 11h0h-6.23a4 4 0 0 0-1.49.29l-3.56 1.42a4 4 0 0 1-1.49.29H5.5h0h-1L3 17zm3-4l1.5-5"/><path d="M6 8h8l2 3"/></g>`),
		g.Group(children),
	)
}