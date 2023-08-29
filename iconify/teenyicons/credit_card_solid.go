package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 2A1.5 1.5 0 0 1 15 3.5V5H0V3.5A1.5 1.5 0 0 1 1.5 2h12Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6v5.5A1.5 1.5 0 0 0 1.5 13h12a1.5 1.5 0 0 0 1.5-1.5V6H0Zm2 4h6V9H2v1Zm11 0h-3V9h3v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}