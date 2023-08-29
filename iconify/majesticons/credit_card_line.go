package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 7a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7zm3-1a1 1 0 0 0-1 1v1h16V7a1 1 0 0 0-1-1H5zm15 4H4v7a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-7zM6 13a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1z"/></g>`),
		g.Group(children),
	)
}