package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nephrology(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21v-4.175q-.25.1-.488.138T8 17q-2.5 0-4.25-1.75T2 11V9q0-2.5 1.75-4.25T8 3q1.25 0 2.125.875T11 6q0 1.25-.875 2.125T8 9v2q1.25 0 2.125.875T11 14v7H9Zm4 0v-7q0-1.25.875-2.125T16 11V9q-1.25 0-2.125-.875T13 6q0-1.25.875-2.125T16 3q2.5 0 4.25 1.75T22 9v2q0 2.5-1.75 4.25T16 17q-.275 0-.513-.038T15 16.825V21h-2Z"/>`),
		g.Group(children),
	)
}