package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxPlotOutlineAlerted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 29a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h17.958l-1.172 2H4v22h28V15.357h2Z" class="clr-i-outline--alerted clr-i-outline-path-1--alerted"/><path fill="currentColor" d="M7 12h10v14H7V12Zm1.6 12.4h6.8v-5.6H8.6v5.6Zm6.8-10.8H8.6v3.6h6.8v-3.6Z" class="clr-i-outline--alerted clr-i-outline-path-2--alerted"/><path fill="currentColor" d="M29 24H19v-8.655c.021.003 1.6.015 1.6.015v1.84h6.8v-1.84l1.6-.003V24Zm-1.6-5.2h-6.8v3.6h6.8v-3.6Z" class="clr-i-outline--alerted clr-i-outline-path-3--alerted"/><path fill="currentColor" d="m26.854 1.144l-5.72 9.86a1.28 1.28 0 0 0 1.1 2h11.45a1.28 1.28 0 0 0 1.1-2l-5.72-9.86a1.28 1.28 0 0 0-2.21 0Z" class="clr-i-outline--alerted clr-i-outline-path-4--alerted clr-i-alert"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}