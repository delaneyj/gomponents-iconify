package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VehicleBusTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.75 5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5ZM4 5.75A3.75 3.75 0 0 1 7.75 2h8.5A3.75 3.75 0 0 1 20 5.75V9.5h1.227a.75.75 0 0 1 0 1.5H20v8.75a1.75 1.75 0 0 1-1.75 1.75h-1.5A1.75 1.75 0 0 1 15 19.75V18.5H9v1.25a1.75 1.75 0 0 1-1.75 1.75h-1.5A1.75 1.75 0 0 1 4 19.75V11H2.75a.75.75 0 0 1 0-1.5H4V5.75ZM16.5 18.5v1.25c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25V18.5h-2Zm-11 0v1.25c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25V18.5h-2Zm2.25-15A2.25 2.25 0 0 0 5.5 5.75V12h13V5.75a2.25 2.25 0 0 0-2.25-2.25h-8.5ZM9 15a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm7 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}