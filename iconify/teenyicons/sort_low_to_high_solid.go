package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortLowToHighSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L5 1.707V15H4V1.707L1.854 3.854l-.708-.708l3-3ZM11 1h-1V0h1.5a.5.5 0 0 1 .5.5V6h1v1h-3V6h1V1Zm-2 9a2 2 0 0 1 2-2h.578a1.885 1.885 0 0 1 1.789 2.482l-1.393 4.176l-.948-.316l.78-2.342H11a2 2 0 0 1-2-2Zm3.14 1l.278-.835A.885.885 0 0 0 11.578 9H11a1 1 0 1 0 0 2h1.14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}