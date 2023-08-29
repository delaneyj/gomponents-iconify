package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 85v86h85v42h-85v86H85v-86H0v-42h85V85h43zm213 235h-42V93l-64 22V79l100-36h6v277z"/>`),
		g.Group(children),
	)
}