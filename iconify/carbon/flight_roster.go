package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightRoster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M26 6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h8v-2H8V6h16v6h2Z"/><path fill="currentColor" d="M10 18h6v2h-6zm0-4h12v2H10z"/><path fill="currentColor" fill-rule="evenodd" d="M22 10v2H10v-2zm3 13l5 2v-2l-5-2.5V18a1 1 0 0 0-2 0v2.5L18 23v2l5-2v3.5L21 28v1l3-1l3 1v-1l-2-1.5z"/>`),
		g.Group(children),
	)
}