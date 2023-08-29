package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundTransferDiagonalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.207 4.788A9.25 9.25 0 0 0 16.525 20.07l-5.897-6.194V16.5a.75.75 0 0 1-1.5 0V12a.75.75 0 0 1 1.294-.517l7.363 7.735A9.25 9.25 0 0 0 7.462 3.938l5.666 6.143V7.5a.75.75 0 0 1 1.5 0V12a.75.75 0 0 1-1.3.508l-7.121-7.72ZM1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}