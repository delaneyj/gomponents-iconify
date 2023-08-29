package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h9m-9 0l-2-2H2l2 2l-2 2h4l2-2Zm9 0l-2-2m2 2l-2 2m1 8.5c2.761 0 5-4.701 5-10.5S18.761 1.5 16 1.5S11 6.201 11 12s2.239 10.5 5 10.5Z"/>`),
		g.Group(children),
	)
}