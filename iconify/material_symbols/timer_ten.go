package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16h3V8h-3v8Zm0 3q-1.25 0-2.125-.875T11 16V8q0-1.25.875-2.125T14 5h3q1.25 0 2.125.875T20 8v8q0 1.25-.875 2.125T17 19h-3Zm-8 0V8H4V5h5v14H6Z"/>`),
		g.Group(children),
	)
}