package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CactusLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 210h-50v-28h6a66.08 66.08 0 0 0 66-66a26 26 0 0 0-52 0a14 14 0 0 1-14 14h-6V56a38 38 0 0 0-76 0v34h-6a14 14 0 0 1-14-14a26 26 0 0 0-52 0a66.08 66.08 0 0 0 66 66h6v68H40a6 6 0 0 0 0 12h176a6 6 0 0 0 0-12ZM96 130H84a54.06 54.06 0 0 1-54-54a14 14 0 0 1 28 0a26 26 0 0 0 26 26h12a6 6 0 0 0 6-6V56a26 26 0 0 1 52 0v80a6 6 0 0 0 6 6h12a26 26 0 0 0 26-26a14 14 0 0 1 28 0a54.06 54.06 0 0 1-54 54h-12a6 6 0 0 0-6 6v34h-52v-74a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}