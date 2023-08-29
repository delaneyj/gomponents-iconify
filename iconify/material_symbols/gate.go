package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19V9q0-1.65 1.175-2.825T9 5h2v6H9v2h2v6H5Zm8 0v-6h2v-2h-2V5h2q1.65 0 2.825 1.175T19 9v10h-6ZM2 17V7h2v10H2Zm18 0V7h2v10h-2Z"/>`),
		g.Group(children),
	)
}