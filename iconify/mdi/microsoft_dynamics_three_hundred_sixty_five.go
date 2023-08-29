package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftDynamicsThreeHundredSixtyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 2l11.25 6.5l-3.75 3L6 8V2m0 7l3.5 2.25L6 22L18 9v6L6 22V9Z"/>`),
		g.Group(children),
	)
}