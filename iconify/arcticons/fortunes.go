package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fortunes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-miterlimit="7" stroke-width=".97"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 18.7h4m-4 4h2.6m-2.6-4v8m6.7 0h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v1.3a2 2 0 0 1-2 2Zm12.7-5.3v3.3a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2v-3.3m0 3.3v2m-12.6-3.3a2 2 0 0 1 2-2h0m-2 0v5.3m24.9-1a1.94 1.94 0 0 1-1.7 1h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.6h-4m-2.3 2.7v-3.3a2 2 0 0 0-2-2h0a2 2 0 0 0-2 2v3.3m0-3.3v-2m-10-1.7v6a.94.94 0 0 0 1 1h.3m-2.3-5.3h2.1"/>`),
		g.Group(children),
	)
}