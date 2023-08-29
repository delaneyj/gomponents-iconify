package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepsLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M246 56a6 6 0 0 1-6 6h-50v42a6 6 0 0 1-6 6h-50v42a6 6 0 0 1-6 6H78v42a6 6 0 0 1-6 6H16a6 6 0 0 1 0-12h50v-42a6 6 0 0 1 6-6h50v-42a6 6 0 0 1 6-6h50V56a6 6 0 0 1 6-6h56a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}