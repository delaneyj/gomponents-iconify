package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.17 8.533C2.21 5.5 4.388 1.5 8 1.5s5.79 4 3.83 7.033L9.592 12H6.408L4.17 8.533ZM5 12.584L2.91 9.347C.305 5.315 3.2 0 8 0s7.694 5.315 5.09 9.347L11 12.584V14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-1.416Zm1.5.916v.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-.5h-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}