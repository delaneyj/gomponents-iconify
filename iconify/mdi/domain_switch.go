package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 9h2v2h-2V9m6 6V5H12V1H2v14h20m-2-2h-8v-2h2V9h-2V7h8v6M8 5V3h2v2H8m0 4V7h2v2H8m0 4v-2h2v2H8M4 5V3h2v2H4m0 4V7h2v2H4m0 4v-2h2v2H4m5 4v2h6v-2l3 3l-3 3v-2H9v2l-3-3l3-3Z"/>`),
		g.Group(children),
	)
}