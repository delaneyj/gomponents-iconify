package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirportShuttleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19q-1.25 0-2.125-.875T3 16H1V5h16l6 6v5h-2q0 1.25-.875 2.125T18 19q-1.25 0-2.125-.875T15 16H9q0 1.25-.875 2.125T6 19Zm9-9h4l-3-3h-1v3Zm-6 0h4V7H9v3Zm-6 0h4V7H3v3Zm3 7.25q.525 0 .888-.363T7.25 16q0-.525-.363-.888T6 14.75q-.525 0-.888.363T4.75 16q0 .525.363.888T6 17.25Zm12 0q.525 0 .888-.363T19.25 16q0-.525-.363-.888T18 14.75q-.525 0-.888.363T16.75 16q0 .525.363.888t.887.362Z"/>`),
		g.Group(children),
	)
}