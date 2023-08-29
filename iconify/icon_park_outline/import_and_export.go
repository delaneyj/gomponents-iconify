package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportAndExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="3.99" d="m14 26l-9 9l9 9m-9-8.992h17.5M34 18l9 9l-9 9m9-8.992H25.5M4.5 24V7.5h39V15"/>`),
		g.Group(children),
	)
}