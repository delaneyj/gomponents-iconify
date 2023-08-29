package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListPlusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M32 72V56a8 8 0 0 1 8-8h176a8 8 0 0 1 8 8v16a8 8 0 0 1-8 8H40a8 8 0 0 1-8-8Zm8 72h176a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8H40a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8Zm112 32H40a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h112a8 8 0 0 0 8-8v-16a8 8 0 0 0-8-8Zm80 8h-16v-16a8 8 0 0 0-16 0v16h-16a8 8 0 0 0 0 16h16v16a8 8 0 0 0 16 0v-16h16a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}