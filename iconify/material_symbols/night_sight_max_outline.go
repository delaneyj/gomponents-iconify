package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightSightMaxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 11V8h-3V6h3V3h2v3h3v2h-3v3h-2Zm-6 8q1.35 0 2.513-.537t1.987-1.488q-3.15-.2-5.325-2.488T8 9q0-.325.025-.638T8.1 7.75q-1.4.775-2.25 2.163T5 13q0 2.5 1.75 4.25T11 19Zm0 2q-3.35 0-5.675-2.325T3 13q0-3.35 2.325-5.675T11 5q.125 0 .25.012t.25.013q-.725.8-1.113 1.812T10 9q0 2.5 1.75 4.25T16 15q.8 0 1.525-.2t1.375-.55q-.45 2.875-2.663 4.813T11 21Zm-.825-6.525Z"/>`),
		g.Group(children),
	)
}