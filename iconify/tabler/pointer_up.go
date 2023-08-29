package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointerUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15.984 13.428l-1.206-1.206l3.113-2.09a1.2 1.2 0 0 0-.309-2.228L4 4l3.904 13.563a1.2 1.2 0 0 0 2.228.308l2.09-3.093l1.217 1.217M19 22v-6m3 3l-3-3l-3 3"/>`),
		g.Group(children),
	)
}