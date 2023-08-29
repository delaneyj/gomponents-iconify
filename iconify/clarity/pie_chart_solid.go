package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 7v22a2 2 0 0 1-2 2s-30-.895-30-2V7a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2Zm-9 12h-8v-8a8 8 0 1 0 8 8Zm-6-2h8a8 8 0 0 0-8-8Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}