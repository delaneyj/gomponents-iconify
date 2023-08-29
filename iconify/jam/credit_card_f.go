package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4H0V2a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2zm0 3v5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V7h20zM4 9a1 1 0 1 0 0 2h1a1 1 0 0 0 0-2H4zm5 0a1 1 0 1 0 0 2h5a1 1 0 0 0 0-2H9z"/>`),
		g.Group(children),
	)
}