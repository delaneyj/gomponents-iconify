package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortHighToLowSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708L4 13.293ZM11 1h-1V0h1.5a.5.5 0 0 1 .5.5V6h1v1h-3V6h1V1Zm-2 9a2 2 0 0 1 2-2h.862a1.68 1.68 0 0 1 1.594 2.212l-1.482 4.446l-.948-.316l.78-2.342H11a2 2 0 0 1-2-2Zm3.14 1l.368-1.104A.68.68 0 0 0 11.862 9H11a1 1 0 1 0 0 2h1.14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}