package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnotherNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 42.5v-7.293l21.386-21.328l7.162 7.248L12.624 42.5ZM38.146 6.245a2.556 2.556 0 0 0-3.607 0l-4.119 4.109l7.214 7.196l4.12-4.11a2.54 2.54 0 0 0 0-3.597Z"/>`),
		g.Group(children),
	)
}