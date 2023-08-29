package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChecklistRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.375 19l-3.55-3.55l1.4-1.4l2.125 2.125l4.25-4.25L22 13.35L16.375 19Zm0-8l-3.55-3.55l1.4-1.4l2.125 2.125l4.25-4.25L22 5.35L16.375 11ZM2 17v-2h9v2H2Zm0-8V7h9v2H2Z"/>`),
		g.Group(children),
	)
}