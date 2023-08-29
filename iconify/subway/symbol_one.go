package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C138.2 0 42.7 95.5 42.7 213.3V512h128V213.3c0-47.1 38.2-85.3 85.3-85.3s85.3 38.2 85.3 85.3V512h128V213.3C469.4 95.5 373.8 0 256 0zM128 469.3H85.4V384H128v85.3zm298.7 0H384V384h42.7v85.3z"/>`),
		g.Group(children),
	)
}