package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PydroidThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34 5v5a4.5 4.5 0 0 0 4 4h5M11.5 29.5v-16h5.2a5.4 5.4 0 1 1 0 10.8h-5.2m16.7 5.2L24 18.9m8 0l-5 14.4a2.34 2.34 0 0 1-2.2 1.6h-.6m7.3-4.7a3.402 3.402 0 0 1 2.5-.7h.8m0 8a2.006 2.006 0 0 0 2-2h0a2.006 2.006 0 0 0-2-2h0a2.006 2.006 0 0 0 2-2h0a2.006 2.006 0 0 0-2-2m-3.3 7.4a3.253 3.253 0 0 0 2.5.7h.8"/>`),
		g.Group(children),
	)
}