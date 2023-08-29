package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardiogramNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCardiogramNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm17 7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3h4a3 3 0 0 1 3 3v31a3 3 0 0 1-3 3H13a3 3 0 0 1-3-3V22.989h8.802l.287-.51l1.325-2.356l1.342 4.696l.541 1.894l1.21-1.555l3.2-4.114h4.51a1 1 0 1 0 0-2h-5.489l-.3.386l-2.29 2.945l-1.459-5.105l-.658-2.305l-1.175 2.089l-2.213 3.935H10V10a3 3 0 0 1 3-3h4Zm3-1a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-8Zm-4 25a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H17a1 1 0 0 1-1-1Zm1 5a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCardiogramNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}