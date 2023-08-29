package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxPlotLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 5H4a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2ZM4 29V7h28v22Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M7 12h10v14H7V12Zm1.6 12.4h6.8v-5.6H8.6v5.6Zm6.8-10.8H8.6v3.6h6.8v-3.6Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M19 24h10V10H19v14Zm1.6-12.4h6.8v5.6h-6.8v-5.6Zm6.8 10.8h-6.8v-3.6h6.8v3.6Z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}