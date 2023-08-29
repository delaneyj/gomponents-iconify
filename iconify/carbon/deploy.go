package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deploy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 2l-6 6l1.415 1.402L22 5.818V28H6V12H4v16a2.002 2.002 0 0 0 2 2h16a2.002 2.002 0 0 0 2-2V5.815l3.586 3.587L29 8Z"/><path fill="currentColor" d="M16 24h-4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2Zm-4-6v4h4v-4Z"/>`),
		g.Group(children),
	)
}