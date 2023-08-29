package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditcardPlusLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9v8a2 2 0 0 0 2 2h6M3 9V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 9h18m0 0v3M7 13h5m6 2v3m0 3v-3m0 0h-3m3 0h3"/>`),
		g.Group(children),
	)
}