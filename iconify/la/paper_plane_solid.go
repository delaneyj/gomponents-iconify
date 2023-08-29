package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m3.594 5.344l.437 1.875L5.97 16l-1.94 8.781l-.437 1.875l1.781-.718l22-9L29.656 16l-2.281-.938l-22-9zm2.781 3.312L21.906 15H7.781zM7.781 17h14.125L6.375 23.344z"/>`),
		g.Group(children),
	)
}