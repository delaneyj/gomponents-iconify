package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paloaltonetworks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.278 15.443l1.705 1.705l-3.426 3.426l-3.427-3.426l8.592-8.591l-1.705-1.705l3.426-3.426l3.427 3.426l-8.592 8.591zM0 12.017l3.426 3.426l8.591-8.59l-3.426-3.427L0 12.017zm11.983 5.13l3.426 3.427L24 11.983l-3.426-3.426l-8.591 8.59z"/>`),
		g.Group(children),
	)
}