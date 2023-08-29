package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinnedOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18M15 4.5l-3.249 3.249m-2.57 1.433L7 10l-1.5 1.5l7 7L14 17l.82-2.186m1.43-2.563L19.5 9M9 15l-4.5 4.5M14.5 4L20 9.5"/>`),
		g.Group(children),
	)
}