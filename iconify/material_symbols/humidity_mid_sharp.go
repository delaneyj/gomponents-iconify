package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumidityMidSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.5q-3.325 0-5.662-2.312Q4 16.875 4 13.6q0-1.65.613-3.063q.612-1.412 1.737-2.487L12 2.5l5.65 5.55q1.125 1.075 1.738 2.487Q20 11.95 20 13.6q0 3.275-2.337 5.588Q15.325 21.5 12 21.5ZM6.025 13h11.95q-.1-1.025-.537-1.9q-.438-.875-1.188-1.6L12 5.3L7.75 9.5q-.75.725-1.188 1.6q-.437.875-.537 1.9Z"/>`),
		g.Group(children),
	)
}