package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRubLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 150a58 58 0 0 0 0-116H88a6 6 0 0 0-6 6v98H56a6 6 0 0 0 0 12h26v20H56a6 6 0 0 0 0 12h26v34a6 6 0 0 0 12 0v-34h50a6 6 0 0 0 0-12H94v-20ZM94 46h54a46 46 0 0 1 0 92H94Z"/>`),
		g.Group(children),
	)
}