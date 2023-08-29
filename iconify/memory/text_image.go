package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h10v10H2V4m2 2v6h6V6H4m10-2h6v2h-6V4m0 4h6v2h-6V8m0 4h6v2h-6v-2M2 16h16v2H2v-2Z"/>`),
		g.Group(children),
	)
}