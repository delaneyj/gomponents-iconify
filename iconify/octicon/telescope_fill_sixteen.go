package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelescopeFillSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.905.42a1.5 1.5 0 0 1 2.144.49l1.692 2.93a1.5 1.5 0 0 1-.649 2.102L2.895 11.815a1.5 1.5 0 0 1-1.95-.602l-.68-1.176a1.5 1.5 0 0 1 .455-1.99L11.905.422Zm-3.374 9.79a.75.75 0 0 1 .944.253l2.644 3.864a.751.751 0 0 1-1.238.847L9 12.424v2.826a.75.75 0 0 1-1.5 0v-2.826l-1.881 2.75a.75.75 0 1 1-1.238-.848l2.048-2.992a.752.752 0 0 1 .293-.252l1.81-.871Zm2.476-3.965v-.001l1.356-.653l-1.52-2.631l-1.243.848ZM3.279 8.119l.835 1.445l1.355-.653l-.947-1.64Z"/>`),
		g.Group(children),
	)
}