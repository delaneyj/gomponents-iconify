package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineDigitalGate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 23h-2v-2a3.003 3.003 0 0 0-3-3H9a3.003 3.003 0 0 0-3 3v2H4v-2a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zM11 6a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5zM2 26h28v2H2zm29.966-14.259a1 1 0 0 0-1.225-.707l-3.086.826L24 7l-1 .268l1.458 5.448l-2.766.74L20 11.206l-1 .268l.934 3.488a1 1 0 0 0 1.224.708l10.1-2.704a1 1 0 0 0 .708-1.225z"/>`),
		g.Group(children),
	)
}