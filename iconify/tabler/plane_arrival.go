package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaneArrival(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15.157 11.81l4.83 1.295a2 2 0 1 1-1.036 3.863L4.462 13.086L3.117 6.514l2.898.776l1.414 2.45l2.898.776l-.12-7.279l2.898.777l2.052 7.797zM3 21h18"/>`),
		g.Group(children),
	)
}