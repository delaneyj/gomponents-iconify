package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoAffix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.918 5.5H7.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H28.082M42.5 38.027L27.653 25.879m-3.67 6.191L36.7 42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.238 42.5l15.679-12.884a2.142 2.142 0 0 0 .788-1.662v-2.197l-.04-12.538h7.496L24.006 1.172L12.85 13.219h7.496v12.605L5.5 37.97"/>`),
		g.Group(children),
	)
}