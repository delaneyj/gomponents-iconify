package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BniMobileBanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.5 18.303v11.394m-10.588 0V18.303l7.548 11.394V18.303M18.2 24a2.848 2.848 0 1 1 0 5.697h-4.7V18.303h4.7a2.848 2.848 0 1 1 0 5.697h0Zm0 0h-4.7"/>`),
		g.Group(children),
	)
}