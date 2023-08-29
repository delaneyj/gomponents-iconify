package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailBulk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 5v10H2v13h17v-4h11V11h-7V5H6zm2 2h13v4h-9v4H8V7zm6 6h14v9h-9v-7h-5v-2zm9 2v2h3v-2h-3zM4 17h13v1.113l-6.04 3.754a.877.877 0 0 1-.92 0L4 18.115V17zm13 3.469V26H4v-5.53l4.982 3.096c.468.291.994.438 1.518.438s1.049-.147 1.516-.44L17 20.47z"/>`),
		g.Group(children),
	)
}