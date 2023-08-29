package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21h18M6 21V6m1.18-2.825C7.43 3.063 7.708 3 8 3h8a2 2 0 0 1 2 2v9m0 4v3M3 3l18 18"/>`),
		g.Group(children),
	)
}