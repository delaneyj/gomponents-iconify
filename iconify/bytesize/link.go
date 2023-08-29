package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 8s6-6 9-3s2 7-3 11s-8 5-10 1m0 7s-6 6-9 3s-2-7 3-11s8-5 10-1"/>`),
		g.Group(children),
	)
}