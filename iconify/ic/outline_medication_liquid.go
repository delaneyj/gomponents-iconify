package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineMedicationLiquid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h12v2H3zM2 21h14V6H2v15zm3-9h2.5V9.5h3V12H13v3h-2.5v2.5h-3V15H5v-3zm15-6c-1.68 0-3 1.76-3 4c0 1.77.83 3.22 2 3.76V21h2v-7.24c1.17-.54 2-1.99 2-3.76c0-2.24-1.32-4-3-4z"/>`),
		g.Group(children),
	)
}