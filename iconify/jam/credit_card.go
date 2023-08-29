package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4V2H2v2h16zm0 3H2v5h16V7zM2 0h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2zm2 9h1a1 1 0 1 1 0 2H4a1 1 0 0 1 0-2zm5 0h5a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2z"/>`),
		g.Group(children),
	)
}