package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightSightAutoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19q1.3 0 2.475-.525t2.025-1.5q-3.2-.2-5.35-2.488T8 9q0-.325.025-.638T8.1 7.75q-1.425.8-2.263 2.2T5 13q0 2.5 1.75 4.25T11 19Zm0 2q-3.35 0-5.675-2.325T3 13q0-3.35 2.325-5.675T11 5q.125 0 .25.012t.25.013q-.725.8-1.113 1.825T10 9q0 2.5 1.75 4.25T16 15q.775 0 1.513-.188t1.387-.562q-.45 2.95-2.7 4.85T11 21Zm2.8-10L17 2h2l3.2 9h-1.9l-.7-2h-3.2l-.7 2h-1.9Zm3.05-3.35h2.3L18 4l-1.15 3.65Zm-6.675 6.825Z"/>`),
		g.Group(children),
	)
}